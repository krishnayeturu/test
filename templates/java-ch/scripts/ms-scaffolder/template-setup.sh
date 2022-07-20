#!/bin/bash
set -e
SCRIPT_DIR="$(dirname -- "$(readlink -f -- "$BASH_SOURCE")")"
PROJECT_DIR="$SCRIPT_DIR/../.."

# Load env.
if [ ! -f "$PROJECT_DIR/.env" ]; then
    cp $PROJECT_DIR/.env.template $PROJECT_DIR/.env
fi
source $PROJECT_DIR/.env

# Rename application src directory.
mv $PROJECT_DIR/src/main/java/com/_2ndwatch/application $PROJECT_DIR/src/main/java/com/_2ndwatch/$PROJECT_SLUG_NO_HYPHEN

