pipeline {
    agent any
    environment {
        IMAGE_REPO_NAME="go-bponline"
        APP_NAME="go-bponline"
        APP_NAMESPACE="backend"
        ENV="main"
    }
    stages {
        stage('Cloning Git') {
            steps {
                script {
                    checkout scm
                }
            }
        }

        stage('Code Scan'){
            when {
                anyOf {
                    branch '${ENV}'
                    expression { 
                        return buildingTag() 
                    }
                }
            }
            environment {
                scannerHome = tool 'sonar-scanner'
            }
            steps {

                withSonarQubeEnv('sonar-staging'){
                    sh "${scannerHome}/bin/sonar-scanner -Dsonar.projectKey=${APP_NAME} -Dsonar.sources=. -Dsonar.exclusions=**/*_test.go -Dsonar.languages=go"
                }
            }
        }
    }
}