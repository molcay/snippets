# 
# Bump version with base semantic versioning

BUMP_TYPE=$(echo "$1" | tr '[:upper:]' '[:lower:]')
NEW_VERSION=""

_bumpVersion() {
  previousVersion=$(<$1)
  majorVersion=$(echo ${previousVersion}|cut -d. -f1)
  minorVersion=$(echo ${previousVersion}|cut -d. -f2)
  patchVersion=$(echo ${previousVersion}|cut -d. -f3)

  case $2 in
    'major')
      majorVersion=$((majorVersion + 1))
      minorVersion=0
      patchVersion=0
      ;;
    'minor')
      minorVersion=$((minorVersion + 1))
      patchVersion=0
      ;;
    *)
      patchVersion=$((patchVersion + 1))
      ;;
  esac

  NEW_VERSION="${majorVersion}.${minorVersion}.${patchVersion}"
}

_bumpVersion $@

echo ${NEW_VERSION}
