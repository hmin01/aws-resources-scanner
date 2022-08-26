#!/bin/sh
# 스크립트 파일 절대 경로
SHELL_PATH="$( cd "$( dirname "$0" )" && pwd -P )"
# Build
echo "Building..."
GOOS=linux CGO_ENABLED=0 go build -o "$SHELL_PATH/bin/scanner"
# Archiving
echo "Archiving..."
cd "$SHELL_PATH/bin"
zip -qq function.zip scanner