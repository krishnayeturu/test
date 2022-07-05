#!/bin/bash
set -e
source .env

# Rename application src directory.
mv src/main/java/com/_2ndwatch/application src/main/java/com/_2ndwatch/$PROJECT_SLUG_NO_HYPHEN

