#!/bin/bash

set -e

IMAGE=${IMAGE:-vugu/wasm-test-suite\:latest}

# Go build for linux
GOOS=linux GOARCH=amd64 go build -o wasm-test-suite-srv .

# Docker build and tag image
docker build -t "${IMAGE}" .

echo "Run with: docker run -ti --rm -p 9222:9222 -p 8846:8846 ${IMAGE}"
echo "To push to DockerHub, run 'docker login' and then 'docker push ${IMAGE}'"
