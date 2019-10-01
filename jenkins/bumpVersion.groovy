stage('Bump Version') {
  steps {
    fileExists 'VERSION'
    script {
      env.WORKSPACE = pwd()
      def version = readFile "${env.WORKSPACE}/VERSION"
      def bumpType  = "${env.bump_type}".toLowerCase()
                
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
    }
  }
}
