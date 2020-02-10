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
        stage('unit-test') {
            steps {
                script {
                    sh 'make ci-unit-test'
                }
            }
        }
        stage('integration-test') {
            steps {
                script {
                    sh 'make ci-integration-test'
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
                        image.push("latest")
                    }
                }
            }
        }
        stage('depoy') {
            steps {
                script {
                    sh 'make ci-deploy'
                }
            }
        }
    }
}

