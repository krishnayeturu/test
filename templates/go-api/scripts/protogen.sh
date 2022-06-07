#!/bin/bash
set -e
source .env

PREFIX="gitlab.com/2ndwatch/microservices/apis"

echo "$PREFIX/$PROJECT_SLUG/pkg/pb"

# mkdir -p pkg/pb
# protoc --go_out=pkg/pb --go-grpc_out=pkg/pb --go_opt=module=$PREFIX/$PROJECT_SLUG/pkg/pb --go-grpc_opt=module=$PREFIX/$PROJECT_SLUG/pkg/pb --proto_path=./proto ./proto/*.proto
# protoc --go_out=pkg/pb --go_opt=module=$PREFIX/$PROJECT_SLUG/pkg/pb --proto_path=./proto/type ./proto/type/*.proto
