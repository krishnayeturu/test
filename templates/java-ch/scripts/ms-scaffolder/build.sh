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
. $SCRIPT_DIR/protogen.sh

gradle build -DrootProjectName=$PROJECT_SLUG --no-daemon --gradle-user-home=/app/.gradle