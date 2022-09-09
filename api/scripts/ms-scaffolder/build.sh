# DO NOT EDIT THIS FILE!
# This file is provided and managed by the ms-scaffolder repository. Any 
# modifications will be destroyed upon update. Please make any necessary 
# modifications in the ms-scaffolder repository:
# https://gitlab.com/2ndwatch/microservices/templates/ms-scaffolder

#!/bin/bash
set -e
SCRIPT_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")"
PROJECT_DIR="$SCRIPT_DIR/../.."

# Load env.
if [ ! -f "$PROJECT_DIR/.env" ]; then
    cp $PROJECT_DIR/.env.template $PROJECT_DIR/.env
fi
source $PROJECT_DIR/.env

# Call protogen script
$SCRIPT_DIR/protogen.sh

mkdir -p $PROJECT_DIR/build
CGO_ENABLED=0 go build -o $PROJECT_DIR/build/main $PROJECT_DIR/cmd