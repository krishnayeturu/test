#!/bin/bash
set -e
source .env

SCRIPT_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")"

# Call protogen script
$SCRIPT_DIR/protogen.sh

gradle build -DrootProjectName=$PROJECT_SLUG --no-daemon --gradle-user-home=/app/.gradle