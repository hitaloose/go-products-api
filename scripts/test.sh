#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")

go test ./internal/...