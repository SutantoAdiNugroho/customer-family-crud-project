FROM php:8.2-fpm-alpine

RUN apk add --no-cache \
    nginx \
    composer \
    git \
    nodejs \
    npm

RUN docker-php-ext-install pdo pdo_mysql opcache session fileinfo tokenizer dom

RUN adduser -D -g '' appuser

WORKDIR /var/www/html

COPY composer.json composer.lock ./

RUN composer install --no-dev --optimize-autoloader

COPY . .

RUN if [ ! -f .env ]; then cp .env.example .env; fi
RUN php artisan key:generate

RUN chown -R appuser:appuser storage bootstrap/cache
RUN chmod -R 775 storage bootstrap/cache

COPY docker/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["sh", "-c", "nginx && php-fpm"]