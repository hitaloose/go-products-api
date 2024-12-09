#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")

cd "$SCRIPT_DIR/.."

swag init -g ./cmd/main.go -o cmd/docs