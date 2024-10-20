pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                // Clones the repository from GitHub and uses the main branch
                git branch: 'main', url: 'https://github.com/anuragtangri/whatsapp.git'
            }
        }
        stage('Install Dependencies') {
            steps {
                // Installs Go dependencies
                // sh 'go mod tidy'
                echo 'Deploying to production...'
            }
        }
        stage('Build') {
            steps {
                // Builds the Go application
                // sh 'go build -o your-app-name'
                echo 'Deploying to production...'
            }
        }
        stage('Test') {
            steps {
                // Runs the Go tests
                // sh 'go test ./...'
                echo 'Deploying to production...'
            }
        }
        stage('Deploy') {
            steps {
                // Add deployment commands here
                echo 'Deploy stage - customize as needed'
            }
        }
    }
}
