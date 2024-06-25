FROM php:8.1-fpm
LABEL authors="tomut0"

# Set working directory
WORKDIR /var/www/html

# Install system dependencies
RUN apt-get update && apt-get install -y \
    vim \
    netcat-traditional \
    cron \
    git \
    curl \
    openssl \
    libonig-dev \
    libicu-dev \
    zip \
    unzip

# Clear cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Install PHP extensions
RUN docker-php-ext-install pdo pdo_mysql mbstring intl

# Copy existing application directory contents
COPY . /var/www/html

# Install Composer
COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

# Install Node.js and npm (version 20)
RUN curl -sL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get update && apt-get install -y nodejs

# Update npm to the latest version
RUN npm install -g npm@latest

# Install all PHP dependencies
RUN composer install --no-dev --optimize-autoloader

# Install npm dependencies and build assets
RUN npm install && npm run build

# Copy existing application directory permissions
RUN chown -R www-data:www-data /var/www/html && chmod -R 755 /var/www/html

# Add crontab file in the cron directory
RUN PHP_PATH=$(which php) && \
    echo "* * * * * $PHP_PATH /var/www/html/artisan schedule:run >> /dev/null 2>&1" > /etc/cron.d/laravel-scheduler

# Give execution rights to the cron job file
RUN chmod 0644 /etc/cron.d/laravel-scheduler

# Apply cron job
RUN crontab /etc/cron.d/laravel-scheduler

# Ensure cron is running
RUN touch /var/log/cron.log

# Expose port 9000 for php-fpm server
EXPOSE 9000

# Entrypoint to handle runtime setup
RUN chmod +x /var/www/html/docker/init.sh
ENTRYPOINT ["/var/www/html/docker/init.sh"]
