# Customer Family CRUD Project

## This project consists Laravel frontend and Go backend with a PostgreSQL database, part of technical test of Bookingtogo

## How to run this project?

### 1. Using Docker Compose
We can get started with Docker Compose. This method sets up all three services, Laravel, Go, and PostgreSQL.

### Prerequisites
* Docker and Docker Compose installed

### Setup 
* Navigate to the root directory of project
* Run the following command to build and start all containers
```console
docker-compose up --build
```
* Once all services are running, we can access applications
    1. Frontend (Laravel): http://localhost:8000
    2. Backend (Go): http://localhost:8080

### 2. Manual setup
We can also run these project without Docker Compose

### Prerequisites
* PHP version 8.2 or higher
* Composer
* Node.js & npm
* Go version 1.23.4 or higher
* PostgreSQL database server running

### A. Backend GO
* Navigate to Go backend directory
* Install Go dependencies
```console
go mod tidy
```
* Set database connection string in the .env file
* Run the application
```console
go run main.go
```
The API will be available at http://localhost:8080

### B. Frontend Laravel
* Navigate to Laravel frontend directory
* Install Composer dependencies:
```console
composer install
```
* Install NPM and build assets
```console
npm install
npm run dev
```
* Copy .env.example to .env and configure Go API URL
* Generate application key
```console
php artisan key:generate
```
* Start the Laravel dev server
```console
php artisan serve --port=8000
```
The frontend will be available at http://localhost:8000
