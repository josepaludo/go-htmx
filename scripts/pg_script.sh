#!/bin/bash

source .env

docker run --rm \
    -e POSTGRES_USER="$DB_USER" \
    -e POSTGRES_PASSWORD="$DB_PASS" \
    -e POSTGRES_DB="$DB_NAME" \
    -p "$DB_PORT":5432 \
    -v $PWD/volumes/db:/var/lib/postgresql/data \
    postgres
