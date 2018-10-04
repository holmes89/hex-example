pipeline {
  agent none
  environment {
    tag = sh(returnStdout: true, script: "git tag --sort version:refname | tail -1").trim()
  }
  stages {
    stage ('Get Code') {
      agent any
      steps {
        git 'https://github.com/Holmes89/hex-example.git'
      }
    }
    stage ('Test') {
      agent{
        docker {
            image 'golang:stretch'
            args '-e XDG_CACHE_HOME=/tmp/.cache'
        }
      }
      steps {
          script {
            sh 'go get -t ./...'
            sh 'go test ./...'
          }
      }
    }
    stage ('Container') {
        agent any
        steps {
          script {
            app = docker.build("holmes89/hex-example")
            docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
                app.push("${GIT_COMMIT}")
                app.push("latest")
            }
          }
        }
    }
    stage ('Tag Container') {
      when { buildingTag() }
      agent any
      steps {
        script {
          app = docker.build("holmes89/hex-example")
          docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
              app.push(tag)
          }
        }
      }
    }
  }
}
