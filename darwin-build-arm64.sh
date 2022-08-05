#!/bin/bash

set -euxo pipefail

GOARCH=${GOARCH:-arm64}

CGO_ENABLED=1 GOOS=darwin GOARCH=$GOARCH go build -o go-m1temperatures -a .