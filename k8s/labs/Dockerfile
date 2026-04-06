FROM php:8.3-apache
WORKDIR /var/www/html
COPY src/ /var/www/html

RUN a2enmod rewrite

EXPOSE 80