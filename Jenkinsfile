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
                echo 'Copying project files to $HOME/src'
                sh 'mkdir -p $HOME/go/src/github.com/emostafa/garson'
                sh 'cp -r * $HOME/go/src/github.com/emostafa/garson'
                sh 'cd $HOME/go/src/github.com/emostafa/garson'
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
