#!/bin/bash
set -e
SCRIPT_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")"
PROJECT_DIR="$SCRIPT_DIR/../.."

mkdir -p $PROJECT_DIR/src/generated/java
protoc --java_out=$PROJECT_DIR/src/generated/java --proto_path=$PROJECT_DIR/src/protobuf/proto $PROJECT_DIR/src/protobuf/proto/*.proto
protoc --java_out=$PROJECT_DIR/src/generated/java --proto_path=$PROJECT_DIR/src/protobuf/proto/type $PROJECT_DIR/src/protobuf/proto/type/*.proto
