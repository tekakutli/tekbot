#!/usr/bin/env sh


cd ./examples/commands/
go run main.go -config dev.toml >> /dev/null 2>&1 &


LLAMA_SAAS_PATH="/home/tekakutli/files/code/llama/llama-saas/"
cd "$LLAMA_SAAS_PATH"
go build
./server >> /dev/null 2>&1 &
