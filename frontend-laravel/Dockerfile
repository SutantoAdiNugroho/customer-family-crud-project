FROM php:8.2-fpm-alpine

RUN apk add --no-cache \
    nginx \
    git \
    nodejs \
    npm \
    php82-tokenizer \
    && docker-php-ext-install \
    pdo_mysql \
    opcache \
    session \
    fileinfo \
    dom

RUN adduser -D -g '' appuser

WORKDIR /var/www/html

COPY . .

RUN composer install --no-dev --optimize-autoloader

RUN if [ ! -f .env ]; then cp .env.example .env; fi
RUN php artisan key:generate

RUN chown -R appuser:appuser storage bootstrap/cache
RUN chmod -R 775 storage bootstrap/cache

COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["sh", "-c", "nginx && php-fpm"]