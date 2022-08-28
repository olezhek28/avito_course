#!/bin/bash

export MIGRATION_DIR=./migrations
export DB_HOST="db"
export DB_PORT="5432"
export DB_NAME="event-service"
export DB_USER="event-service-user"
export DB_PASSWORD="event-service-password"
export DB_SSL=disable

export PG_DSN="host=${DB_HOST} port=${DB_PORT} dbname=${DB_NAME} user=${DB_USER} password=${DB_PASSWORD} sslmode=${DB_SSL}"

sleep 2 && goose -dir ${MIGRATION_DIR} postgres "${PG_DSN}" up -v