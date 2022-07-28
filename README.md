# gofound-blogger
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
./bin/blog-service
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

# PUT update the post
BODY='{"title":"updated test","content":"updated test content","author":"bartb","img_urls":["https://img.nl/98", "https://img.nl/99"]}'
curl -X PUT -d "$BODY" localhost:4000/v1/posts/1

# GET updates
curl -X GET http://localhost:4000/v1/posts/1

# test delete
curl -X DELETE localhost:4000/v1/posts/1


```
