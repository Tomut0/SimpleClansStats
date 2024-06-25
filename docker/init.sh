#!/bin/bash
set -e

# Function to check if a TCP connection can be made
function check_tcp_connection {
  nc -z -w5 "$DB_HOST" 3306
}

# Check if the .env file exists
if [ ! -f .env ]; then
  echo ".env file not found!"
  exit 1
fi

# Wait for the database to be ready
echo "Waiting for the database to be ready..."
until check_tcp_connection; do
  echo "Database is unavailable - sleeping"
  sleep 3
done

echo "Database is ready - continuing"

# Extract the APP_KEY value from the .env file
APP_KEY=$(grep '^APP_KEY=' .env | cut -d '=' -f2)

# Check if APP_KEY is empty or not set
if [ -z "$APP_KEY" ]; then
  echo "Generating Laravel application key..."
  php artisan key:generate --force

  # Extract new APP_KEY
  NEW_APP_KEY=$(grep '^APP_KEY=' .env | cut -d '=' -f2)
  export APP_KEY=$NEW_APP_KEY
else
  echo "Laravel application key is already set: $APP_KEY"
fi

# Run database migrations
echo "Running database migrations..."
php artisan migrate --force

# Run cron jobs
cron

# Start PHP-FPM
php-fpm

