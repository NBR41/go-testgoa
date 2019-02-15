#!/bin/sh

while [ 1 ]; do
    nc -z ${MYSQL_HOST} 3306 1>/dev/null 2>&1
    if [ $? != 0 ]; then
        echo "no database"
        sleep 5
        continue
    fi
    break
done

exec "/go/bin/myinventory" "env=docker" "db_host=${MYSQL_HOST}" "db_user=${MYSQL_USER}" "db_password=${MYSQL_PASSWORD}" "db_name=${MYSQL_DBNAME}"
