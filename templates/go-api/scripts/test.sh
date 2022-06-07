#!/bin/bash
set -e

go test --test.short -race ./... $@