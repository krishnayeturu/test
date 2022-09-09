#!/bin/bash
set -e
SCRIPT_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")"
PROJECT_DIR="$SCRIPT_DIR/../.."

# Load env.
if [ ! -f "$PROJECT_DIR/.env" ]; then
    cp $PROJECT_DIR/.env.template $PROJECT_DIR/.env
fi
source $PROJECT_DIR/.env

# Define variables and create output directory.
MODULE="gitlab.com/2ndwatch/microservices/integrations/$PROJECT_SLUG/pkg/pb"
OUTPUT_DIR="$PROJECT_DIR/pkg/pb"
PROTO_DIR="$PROJECT_DIR/proto"
mkdir -p $OUTPUT_DIR

# Generate root protos.
protoc \
    --go_out=$OUTPUT_DIR \
    --go-grpc_out=$OUTPUT_DIR \
    --go_opt=module=$MODULE \
    --go-grpc_opt=module=$MODULE \
    --proto_path=$PROTO_DIR $PROTO_DIR/*.proto

# Generate type protos.
protoc \
    --go_out=$OUTPUT_DIR \
    --go_opt=module=$MODULE \
    --proto_path=$PROTO_DIR/type $PROTO_DIR/type/*.proto