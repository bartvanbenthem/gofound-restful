# gofound-restful
Project template for creating an RESTfull webservice in Go, with relational database backend and example content.


### Used packages
* github.com/go-chi/chi 
* github.com/jackc/pgconn
* github.com/jackc/pgx/v4
* github.com/dgrijalva/jwt-go
* golang.org/x/crypto

## Prerequisites
Run db script to install, configure and start PostgreSQL, when the database is started this scrpt creates the database and inserts data from the go_software.sql configuration file.
```shell
$ sudo ./build/scripts/db/install-postgresql.sh
```
Test the repository
``` shell
$ go test -v -cover ./...
```

## Start the API Server
```shell
$ ./build/bin/server
```
```bash
2021/01/01 00:00:00 Connecting to database...
2021/01/01 00:00:00 Connected to database
2021/01/01 00:00:00 Starting server on port 4000
```

## Test the API
```shell
$ ./build/client-test/client-test
```
