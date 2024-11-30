#!/usr/bin/env bash
export DBUSER=root
export DBPASS=local-dev

echo "___ Starting mariadb ___"
docker-compose up -d

echo "___ Generating mysql data ___"
mysql -u "$DBUSER" -P 3306 -h 127.0.0.1 -p"$DBPASS" < data.sql

echo "___ Running main.go ___"
go run ../main.go