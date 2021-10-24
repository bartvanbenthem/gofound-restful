#!/bin/bash

# go test -v -cover cmd/api/*
# go test -v -cover internal/handlers/*

# go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

go build -o build/bin/server $(ls -1 cmd/api/*.go | grep -v _test.go)
./build/bin/server