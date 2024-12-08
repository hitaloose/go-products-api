#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")

export GIN_MODE=debug

gow run "$SCRIPT_DIR/../cmd/main.go"