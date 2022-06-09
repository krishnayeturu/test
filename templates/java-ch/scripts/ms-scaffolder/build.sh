#!/bin/bash
set -e
source .env

gradle build -DrootProjectName=$PROJECT_SLUG --no-daemon --gradle-user-home=/app/.gradle