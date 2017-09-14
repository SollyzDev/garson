pipeline {
    agent {
        docker { image 'golang:alpine' }
    }

    environment {
        GOHOME = "${HOME}"
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building Garson....'
                echo '${HOME}'
                sh 'export GOPATH=$HOME'
                sh 'go get'
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
