#!/bin/sh

set -e

echo "Running database migrations..."

goose \
    -dir /app/sql/migrations \
    up

echo "Starting SimpleClansStats..."

exec /app/simpleclansstats "$@"
