# GoFound-Blogger
RESTful Blog webservice, this repo also serves as a project template for RESTful CRUD API in general.

## Start PostgreSQL Container
```bash
cd project
docker-compose up &
cd ..
```

### Configure PostgreSQL Database
```bash
# set DSN with default DB
export BLOG_DB_DSN='postgres://postgres:password@localhost/default_database?sslmode=disable'
psql $BLOGGER_DB_DSN
# create DB
CREATE DATABASE blog;
exit;
#

# set DSN with new DB
export BLOG_DB_DSN='postgres://postgres:password@localhost/blog?sslmode=disable'
# install migrate tool
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate.linux-amd64 $GOPATH/bin/migrate
migrate --version

# create new migration files
# migrate create -seq -ext=.sql -dir=./migrations 'name_file'

# migrate up
migrate -path=./migrations -database=$BLOGGER_DB_DSN up

```

## Build & Run Blog webservice
```bash
# build bin
go build -o bin/blog-service ./cmd/api/
# set DSN
export BLOG_DB_DSN='postgres://postgres:password@localhost/blog?sslmode=disable'
# run blog service
./bin/blog-service --port=4000 \
    --env='development' \
    --db-dsn=$BLOG_DB_DSN \
    --db-max-open-conns=25 \
    --db-max-idle-conns=25 \
    --db-max-idle-time='15m'
```

### example requests
```bash
# GET server healthcheck
curl -X GET http://localhost:4000/v1/healthcheck

# POST Create first post
BODY='{"title":"testing","content":"test content to display","author":"bartb","img_urls":["https://img.nl/01", "https://img.nl/02"]}'
curl -i -d "$BODY" localhost:4000/v1/posts

# POST empty body error
curl -X POST localhost:4000/v1/posts

# GET first post
curl -X GET http://localhost:4000/v1/posts/1

# PATCH full update the post
BODY='{"title":"updated test","content":"updated test content","author":"bartb","img_urls":["https://img.nl/98", "https://img.nl/99"]}'
curl -X PATCH -d "$BODY" localhost:4000/v1/posts/1

# PATCH partial update the post
BODY='{"author":"John"}'
curl -X PATCH -d "$BODY" localhost:4000/v1/posts/1

# PATCH Error on nil value
# JSON items with null values will be ignored and will remain unchanged
BODY='{"title":"","author":"John"}'
curl -X PATCH -d "$BODY" localhost:4000/v1/posts/1

# PATCH Testing for DATA races
xargs -I % -P8 curl -X PATCH -d '{"author": "bartb"}' "localhost:4000/v1/posts/1" < <(printf '%s\n' {1..8})

# GET updates
curl -X GET http://localhost:4000/v1/posts/1

# DELETE post
curl -X DELETE localhost:4000/v1/posts/1

```
