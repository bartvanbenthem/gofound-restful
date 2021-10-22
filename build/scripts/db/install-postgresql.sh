#!/bin/bash

## INSTALLATION POSTGRESQL (FEDORA 33) 

# Install the repository RPM
sudo dnf install -y https://download.postgresql.org/pub/repos/yum/reporpms/F-34-x86_64/pgdg-fedora-repo-latest.noarch.rpm

# Install PostgreSQL
sudo dnf install -y postgresql14-server

# Optionally initialize the database and enable automatic start
sudo /usr/pgsql-14/bin/postgresql-14-setup initdb
sudo systemctl enable postgresql-14
sudo systemctl start postgresql-14

# configure firewalld when running
sudo firewall-cmd --add-service=postgresql --permanent
sudo firewall-cmd --reload

# config file - listen_addresses = '*' or listen_addresses = 'localhost'
sudo cat /var/lib/pgsql/14/data/postgresql.conf | grep listen_addresses

# copy .sql file to postgres user ~
sudo cp build/scripts/db/go_software.sql /var/lib/pgsql/

# alter administrator password
sudo su - postgres 
psql -c "alter user postgres with password 'password'"

# create database
createdb go_software -O postgres
# configure database, 
psql -d go_software -f go_software.sql