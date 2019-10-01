

String readFileString(String filePath) {
    File file = new File(filePath)
    String fileContent = file.text
    return fileContent
}

def version = readFileString(args[0])
def bumpType  = ( (args.length >= 2) ? args[1] : 'patch' ).toLowerCase()

def (majorVersion, minorVersion, patchVersion) = version.split('\\.').collect { it as int }

switch(bumpType) {
  case "major":
    majorVersion += 1
    minorVersion = 0
    patchVersion = 0
    break
  case "minor":
    minorVersion += 1
    patchVersion = 0
    break
  default:
    patchVersion += 1
    break
}

def newVersion = majorVersion + "." + minorVersion + "." + patchVersion
println(newVersion)

