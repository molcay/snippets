
SEARCH_FOR=$1
SEARCH_DIR=$2

searchByFileContent() {
  grep -iRl ${SEARCH_FOR} ${SEARCH_DIR}
}

searchByFileContent
