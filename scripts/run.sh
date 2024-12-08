#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")

export GIN_MODE=release

go run "$SCRIPT_DIR/../cmd/main.go"