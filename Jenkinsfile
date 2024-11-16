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
                git branch: 'main', url: 'https://github.com/ridhokurniawan-u/Home-Automation.git' 
                echo 'BOO!'
                echo 'Run Finished'
            }
        }
    }
}
