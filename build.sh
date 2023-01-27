#!/bin/bash

OS = $(shell uname -s)
ARCH = $(shell uname -m)

mkdir -p build
go build -o build/ ./main.go
cp config.toml build/
tar -czvf favpics-helper-$(OS)-$(ARCH).tar.gz build/