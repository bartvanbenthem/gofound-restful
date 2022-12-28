#!/bin/bash

# build bin
#go build -o bin/blog-service ./cmd/api/
# set DSN
export BLOG_DB_DSN='postgres://postgres:password@localhost/blog?sslmode=disable'
# run blog service
go run ./cmd/api/ --port=4000 \
    --env='development' \
    --db-dsn=$BLOG_DB_DSN \
    --db-max-open-conns=25 \
    --db-max-idle-conns=25 \
    --db-max-idle-time='15m'



