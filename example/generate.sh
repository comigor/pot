#!/bin/bash

OUT_PATH=./handler/pb

if [ -d $OUT_PATH ]; then
  rm -rf $OUT_PATH
fi

mkdir -p $OUT_PATH

protoc --go_out=$OUT_PATH --go_opt=paths=source_relative \
  --plugin=protoc-gen-go-http="$(which protoc-gen-go-http)" \
  --go-http_out=$OUT_PATH --go-http_opt=paths=source_relative \
  -I=./proto ./proto/**/*.proto
