# GoFound-RESTFul
RESTful Blog webservice, this repo also serves as a project template for RESTful CRUD API in general.

## Start PostgreSQL Container
```bash
cd project
docker-compose up -d
cd ..
```

### Configure PostgreSQL Database
```bash
# DATABSE SHOULD ALREADY BE CREATED TROUGH DOCKER COMPOSE 
# OTHERWISE RUN THE FOLLOWING COMMANDS:
# set DSN with default DB
export BLOG_DB_DSN='postgres://postgres:password@localhost/postgres?sslmode=disable'
psql $BLOG_DB_DSN

# psql=#
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
migrate -path=./migrations -database=$BLOG_DB_DSN up

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

# POST Create first posts
BODY='{"title":"testing","content":"test content to display","author":"bartb","img_urls":["https://img.nl/01", "https://img.nl/02"]}'
BODY2='{"title":"atest","content":"information for testing","author":"johnd","img_urls":["https://img.local/02", "https://img.local/03"]}'
BODY3='{"title":"xtest","content":"another test page","author":"alice","img_urls":["https://noimg.nl/noimg"]}'
BODY4='{"title":"ntest","content":"testing content","author":"robert","img_urls":["https://imgages.com/banner", "https://imgages.com/icon"]}'
curl -i -d "$BODY" localhost:4000/v1/posts
curl -i -d "$BODY2" localhost:4000/v1/posts
curl -i -d "$BODY3" localhost:4000/v1/posts
curl -i -d "$BODY4" localhost:4000/v1/posts

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

# LIST all Posts
curl "localhost:4000/v1/posts"
# LIST Posts & TEST query string input
curl "localhost:4000/v1/posts?title=updated+test"
curl "localhost:4000/v1/posts?img_urls=https://img.nl/98"
# TEST query string input ERROR handling
curl "localhost:4000/v1/posts?page=-1&page_size=-1&sort=foo"
# TEST Partial Text search
curl "localhost:4000/v1/posts?title=test"
# TEST SORTING
curl "localhost:4000/v1/posts?sort=-title"
# TEST Page size
curl "localhost:4000/v1/posts?page_size=2"
curl "localhost:4000/v1/posts?page_size=3&page=2"
curl "localhost:4000/v1/posts?page_size=2&page=3"
# TEST too-high page value,
curl "localhost:4000/v1/posts?page=100"

# titles
curl "localhost:4000/v1/posts" 

# DELETE post
curl -X DELETE localhost:4000/v1/posts/1

```
