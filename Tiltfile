# -*- mode: Python -*-

docker_build('executor', '.', dockerfile='Dockerfile.executor')
k8s_yaml('deployments/executor.yaml')
k8s_resource('executor')

docker_build('server', '.', dockerfile='Dockerfile.server')
k8s_yaml('deployments/server.yaml')
k8s_resource('server')