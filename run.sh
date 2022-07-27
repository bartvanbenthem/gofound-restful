#!/bin/bash

export BLOG_DB_DSN='postgres://postgres:password@localhost/blog?sslmode=disable'

go build -o bin/blogger ./cmd/api/

./bin/blogger


