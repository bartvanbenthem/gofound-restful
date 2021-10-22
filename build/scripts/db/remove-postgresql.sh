#!/bin/bash

# DELETE DATABASE
sudo su - postgres 
# postgres@
psql 
# postgres=#
DROP DATABASE go_software;
exit
# postgres@
exit
# user@
sudo rm -fr /var/lib/pgsql
sudo dnf remove -y postgresql14-server