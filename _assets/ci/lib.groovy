def gitCommit() {
  return GIT_COMMIT.take(6)
}

def gitBranch() {
  return env.GIT_BRANCH.replace('origin/', '')
}

@NonCPS
def getVersion(branch, sha, buildNumber) {
  version = branch.replaceAll(/\//, '-')

  if (sha?.trim()) {
    version = version + '-g' + sha
  }
  if (buildNumber?.trim()) {
    version = version + '-' + buildNumber
  }
  return version
}

def deployToArtifactory() {
  // For branch builds, replace the old artifact. For develop keep all of them.
  def version = (gitBranch() == 'develop' ?
    getVersion(gitBranch(), gitCommit(), '') :
    getVersion(gitBranch(), gitCommit(), env.BUILD_ID)
  )
  def server = Artifactory.server 'artifactory'
  def uploadSpec = """{
      "files": [
          {
              "pattern": "build/bin/statusgo-android-16.aar",
              "target": "libs-release-local/status-im/status-go/${version}/status-go-${version}.aar"
          },
          {
              "pattern": "build/bin/statusgo-ios-9.3-framework/status-go-ios.zip",
              "target": "libs-release-local/status-im/status-go-ios-simulator/${version}/status-go-ios-simulator-${version}.zip"
          }
      ]
  }"""

  def buildInfo = Artifactory.newBuildInfo()
  buildInfo.env.capture = false
  buildInfo.name = 'status-go (' + gitBranch() + '-' + gitCommit() + ')'
  server.upload(uploadSpec, buildInfo)
  server.publishBuildInfo(buildInfo)
}

return this
