# Create DB 
export BLOG_DB_DSN='postgres://postgres:password@localhost/default_database?sslmode=disable'
psql $BLOGGER_DB_DSN

#
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
migrate create -seq -ext=.sql -dir=./migrations 'name_file'

# migrate up
migrate -path=./migrations -database=$BLOGGER_DB_DSN up