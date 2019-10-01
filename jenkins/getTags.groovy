
def gettags = ("git ls-remote -t -h https://<USERNAME>:<PASSWORD>@<REPO_ADDRESS>.git").execute()
return gettags.text.readLines().collect {
  it.split()[1].replaceAll("\\^\\{\\}", '')
}.findAll {
  it.contains('tags')
}.collect {
  it.replaceAll('refs/heads/', '').replaceAll('refs/tags/', '')
}.sort(false) { t1, t2 ->
    def t1Parts = t1.split('\\.').collect { it as int }
    def major1, minor1, patch1
    major1 = t1Parts[0]
    minor1 = t1Parts[1]
    patch1 = t1Parts[2]

    def t2Parts = t2.split('\\.').collect { it as int }
    def major2, minor2, patch2
    major2 = t2Parts[0]
    minor2 = t2Parts[1]
    patch2 = t2Parts[2]

    if (major1 == major2)
        if (minor1 == minor2) 
            patch1 <=> patch2
        else 
            minor1 <=> minor2        
    else 
        major1 <=> major2
}.reverse()
