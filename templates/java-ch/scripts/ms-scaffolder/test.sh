#!/bin/bash
set -e
source .env

SCRIPT_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")"

# Call build script.
$SCRIPT_DIR/build.sh

gradle test -DrootProjectName=$PROJECT_SLUG --no-daemon --gradle-user-home=/app/.gradle