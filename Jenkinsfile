pipeline {
    agent any

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
