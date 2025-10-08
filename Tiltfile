# -*- mode: Python -*-

docker_build('executor', '.', dockerfile='Dockerfile.executor')
k8s_yaml('deployments/executor.yaml')
k8s_resource('executor')