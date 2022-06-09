#!/bin/bash
set -e
source .env

gradle test -DrootProjectName=$PROJECT_SLUG --no-daemon --gradle-user-home=/app/.gradle