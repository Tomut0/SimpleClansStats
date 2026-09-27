#!/bin/sh

set -e

echo "Running database migrations..."

goose \
    -dir /app/sql/migrations \
    sqlite3 "$DATABASE_PATH" \
    up

echo "Starting SimpleClansStats..."

exec /app/simpleclansstats