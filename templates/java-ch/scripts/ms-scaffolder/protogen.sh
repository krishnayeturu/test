#!/bin/bash
set -e

mkdir -p ./src/generated/java
protoc --java_out=./src/generated/java --proto_path=./src/protobuf/proto ./src/protobuf/proto/*.proto
protoc --java_out=./src/generated/java --proto_path=./src/protobuf/proto/type ./src/protobuf/proto/type/*.proto
