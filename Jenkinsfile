pipeline {
    environment {
        registryURL = 'https://registry.hub.docker.com'
        registryCredential = 'docker-credential'
        registryRepository = 'thdatcs/go-protobuf'
    }
    agent any
    stages {
        stage('base') {
            steps {
                script {
                    sh 'make ci-base'
                }
            }
        }
        stage('lint') {
            steps {
                script {
                    sh 'make ci-lint'
                }
            }
        }
        stage('test') {
            steps {
                script {
                    sh 'make ci-test'
                }
            }
        }
        stage('build') {
            steps {
                script {
                    sh 'make ci-build'
                }
            }
        }
        stage('publish') {
            steps {
                script {
                    docker.withRegistry(registryURL, registryCredential) {
                        image = docker.image(registryRepository)
                        image.push("${env.BUILD_NUMBER}")
                    }
                }
            }
        }
    }
}

