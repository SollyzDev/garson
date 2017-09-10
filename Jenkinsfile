pipeline {
    agent {
        docker 'go:alpine'
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building Garson....'
                sh 'go build'
            }
        }

        stage('Test') {
            steps {
                echo 'Running Test Cases'
                sh 'go test .'
            }
        }
    }
}
