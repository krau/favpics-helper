#!/bin/bash

OS=$(uname -s)
ARCH=$(uname -m)

mkdir -p build
go build -o build/ ./main.go
cp config.toml build/
tar -czvf favpics-helper-$OS-$ARCH.tar.gz build/
rm -rf build
mkdir -p build
mv favpics-helper-$OS-$ARCH.tar.gz build/