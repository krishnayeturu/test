#!/bin/bash
set -e
source .env

# TODO: Update this to work the way we want it to.
0 GOOS=linux GOARCH=$* go build -o $PROJECT_SLUG $@ ./cmd/PLACEHOLDER_DIRECTORY