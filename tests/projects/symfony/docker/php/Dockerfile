FROM php:8.5-fpm

RUN mkdir -p /app/node_modules \
    && chown -R www-data:www-data /app \
    && chown -R www-data:www-data /var/www

# SET TIME ZONE
ENV TZ="Europe/Zurich"

RUN apt-get update && apt-get install -y --no-install-recommends \
    vim \
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
    libmagickwand-dev \
    imagemagick \
    && rm -rf /var/lib/apt/lists/*

    RUN docker-php-ext-configure gd \
        --with-freetype \
        --with-jpeg \
    && docker-php-ext-install -j$(nproc) \
        pdo \
        pdo_mysql \
        bcmath \
        mbstring \
        gd \
        intl

RUN pecl install imagick \
 && docker-php-ext-enable imagick

RUN pecl install xdebug \
&& docker-php-ext-enable xdebug

# COMPOSER INSTALL
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

# INSTALL NODE.JS 20 LTS
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash - \
    && apt-get install -y nodejs

WORKDIR /app
USER root
