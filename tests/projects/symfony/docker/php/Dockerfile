FROM php:8.5-fpm

RUN mkdir -p /app/node_modules \
    && chown -R www-data:www-data /app \
    && chown -R www-data:www-data /var/www

RUN apt-get update && apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    pkg-config \
    && rm -rf /var/lib/apt/lists/*

RUN apt-get update && apt-get install -y --no-install-recommends \
    libcurl4-openssl-dev \
    libzip-dev \
    zlib1g-dev \
    libpng-dev \
    libjpeg62-turbo-dev \
    libfreetype6-dev \
    libonig-dev \
    libxml2-dev \
    libldap2-dev \
    libicu-dev \
    && rm -rf /var/lib/apt/lists/*

    RUN docker-php-ext-install -j$(nproc) \
        pdo \
        pdo_mysql \
        gd

# COMPOSER INSTALL
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

WORKDIR /app
USER root
