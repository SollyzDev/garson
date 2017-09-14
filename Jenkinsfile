pipeline {
    agent {
        docker { image 'golang:alpine' }
    }

    environment {
        GOPATH = "${HOME}"
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building Garson....'
                echo '${HOME} ${pwd}'
                sh 'export GOPATH=`pwd`'
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
