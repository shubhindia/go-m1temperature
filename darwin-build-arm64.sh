#!/bin/bash

set -euxo pipefail

GOARCH=${GOARCH:-arm64}


CC=arm64-apple-darwin20-cc CXX=arm64-apple-darwin20-c++ MACOSX_DEPLOYMENT_TARGET=10.10.0 CGO_ENABLED=1 GOOS=darwin GOARCH=$GOARCH go build -ldflags "-X main.version=1.0.0" -o go-m1temperature -a .