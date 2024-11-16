pipeline {
    agent any
    triggers {
    githubPush()
    }
    stages {
        stage('Testing Build') {
            steps {
                echo 'Starting Clone Stage'
                echo 'Build Finished'
            }
        }
        stage('Testing Run'){
            steps {
                echo 'testing Run'
                echo 'Run Finished'
            }
        }
    }
}
