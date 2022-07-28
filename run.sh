#!/bin/bash

# build bin
go build -o bin/blog-service ./cmd/api/
# set DSN
export BLOG_DB_DSN='postgres://postgres:password@localhost/blog?sslmode=disable'
# run blog service
./bin/blog-service


