pipeline {
    agent any
    tools {
        go("go-1.9")
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
                  sh("make all")
                }
            }
        }
    }

    post {
        always {
            junit(testResults: "**/build/reports/*.xml", allowEmptyResults: false, keepLongStdio: true)
            archiveArtifacts(artifacts: "**/build/binaries/gw_*", fingerprint: true,  onlyIfSuccessful: true)
        }
    }
}