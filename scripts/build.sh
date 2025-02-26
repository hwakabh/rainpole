#!/bin/sh
set -e

# For files whose name contains whitespaces
IFS=$'\n'

FILES=( $(find ./ -type f -iname "*.go") )
echo "Sources: ${FILES[@]}" >&2

# Build single binaries to ./cmd
go build -o ./cmd/rainpole "${FILES[@]}"
