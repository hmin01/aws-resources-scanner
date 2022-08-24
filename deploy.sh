#!/bin/sh
# 스크립트 파일 절대 경로
SHELL_PATH="$( cd "$( dirname "$0" )" && pwd -P )"
# Deploy
aws lambda update-function-code --function-name AWSResourceScanner --zip-file "fileb://$SHELL_PATH/bin/function.zip" --profile private