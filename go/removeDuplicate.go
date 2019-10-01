package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strings"
)

// 8KB
const fileChunk = 8192

// current working directory
const cwd = "."

func main() {
	var checkSums = make(map[string]bool)
	var stopper bool // dummy variable for user interaction

	fileList := listFiles(cwd)

	fmt.Printf("Total number of files: %d\n", len(fileList))
	fmt.Print("Press enter to start")

	_, _ = fmt.Scanln(&stopper)

	for _, f := range fileList {
		checksum, err := md5Checksum(f)
		if err != nil {
			panic(err)
		}

		checksumStr := fmt.Sprintf("%x", checksum) // convert []byte to string
		if checkSums[checksumStr] {
			fmt.Printf("Found same file: %s\n", f)
			if err := os.Remove(f); err != nil {
				fmt.Printf("cannot remove file: %s", f)
			}
		} else {
			checkSums[checksumStr] = true
		}
	}

	fmt.Printf("Total number of deleted files: %d\n", len(fileList)-len(checkSums))
	fmt.Print("Press enter to close...")
	_, _ = fmt.Scanln(&stopper)
}

// listFiles list all the files in given directory except the files which are ends with ".exe"
func listFiles(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}

	var fileList []string

	for _, f := range files {
		filename := path.Join(dirName, f.Name())

		if strings.HasSuffix(filename, ".exe") || f.IsDir() {
			continue
		} else {
			fileList = append(fileList, filename)
		}
	}

	return fileList
}

// md5Checksum function open the given file, read it and calculate md5 hash for the file
// Source: https://www.socketloop.com/tutorials/how-to-generate-checksum-for-file-in-go
func md5Checksum(path string) ([]byte, error) {
	// Open the file for reading
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cannot find file:", path)
		return []byte{}, fmt.Errorf("cannot find file: %s", path)
	}

	defer file.Close()

	// Get file info
	info, err := file.Stat()
	if err != nil {
		fmt.Println("Cannot access file:")
		return []byte{}, fmt.Errorf("cannot access file: %s", path)
	}

	// Get the file size
	fileSize := info.Size()

	// Calculate the number of blocks
	blocks := uint64(math.Ceil(float64(fileSize) / float64(fileChunk)))

	// Start hash
	hash := md5.New()

	// Check each block
	for i := uint64(0); i < blocks; i++ {
		// Calculate block size
		blockSize := int(math.Min(fileChunk, float64(fileSize-int64(i*fileChunk))))

		// Make a buffer
		buf := make([]byte, blockSize)

		// Make a buffer
		if _, err := file.Read(buf); err != nil {
			return []byte{}, fmt.Errorf("file read to buffer failed: %s", path)
		}

		// Write to the buffer
		if _, err := io.WriteString(hash, string(buf)); err != nil {
			return []byte{}, fmt.Errorf("write buffer to hash failed: %s", path)
		}
	}

	return hash.Sum(nil), nil
}
