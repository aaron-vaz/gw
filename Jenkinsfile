def getRepoURL() {
    return sh(script: "git config --get remote.origin.url", returnStdout: true).trim()
}

def getCommitID() {
    return sh(script: "git rev-parse HEAD", returnStdout: true).trim()
}

pipeline {
    agent any
    tools {
        go("go")
    }

    stages {
        stage("Checkout") {
            steps {
                checkout scm
                sh("git reset --hard")
                sh("git clean -fdx")
            }
        }

        stage("Build") {
            steps {
                script {
                    sh("make ci")
                    currentBuild.displayName = readFile("${env.WORKSPACE}/build/version.txt").trim()
                }
            }
        }
    }

    post {
        always {
            junit(testResults: "**/build/reports/*.xml", allowEmptyResults: false, keepLongStdio: true)
            archiveArtifacts(artifacts: "**/build/binaries/gw_*", fingerprint: true, onlyIfSuccessful: true)
        }

        success {
            script {
                if("${env.BRANCH_NAME}" == "master") {
                    githubRelease(repoURL: getRepoURL(),
                        releaseTag: "v${currentBuild.displayName}",
                        commitish: getCommitID(),
                        releaseName: "Release v${currentBuild.displayName}",
                        releaseBody: "",
                        isPreRelease: false,
                        isDraftRelease: false,
                        artifactPatterns: "**/build/binaries/gw_*")
                }
            }
        }
    }
}