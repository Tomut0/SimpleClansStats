#!/bin/sh

set -e

echo "Running database migrations..."

goose \
    -dir /app/migrations \
    sqlite3 "$DATABASE_PATH" \
    up

echo "Starting SimpleClansStats..."

exec /app/simpleclansstats