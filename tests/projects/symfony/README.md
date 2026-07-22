# Code snippets

![Tailwind](https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-B73BFE?style=for-the-badge&logo=vite&logoColor=FFD62E)
![Node.js](https://img.shields.io/badge/Node%20js-339933?style=for-the-badge&logo=nodedotjs&logoColor=white)<br>
![Symfony](https://img.shields.io/badge/Symfony-000000?style=for-the-badge&logo=Symfony&logoColor=white)
![PHP](https://img.shields.io/badge/PHP-777BB4?style=for-the-badge&logo=php&logoColor=white)
![Mysql](https://img.shields.io/badge/MySQL-005C84?style=for-the-badge&logo=mysql&logoColor=white)

<br>

## Requirements

- [Docker Compose](https://docs.docker.com/compose/install)

<br>

## Setup

Freshly set up everything needed for the project

```bash
docker compose -f docker/docker-compose.yaml up -d
```

## Frontend Commands

```bash
npm i & docker exec -it codesnippets-fpm npm i
```

Run the development servers

```bash
docker exec -it codesnippets-fpm npm run dev
```

Build the frontend

```bash
docker exec -it codesnippets-fpm npm run build
```

<br>

## Backend Commands

Install dependencies

```bash
docker exec -it codesnippets-fpm composer i
```

Migrate the DB structure in the database

```bash
docker exec -it codesnippets-fpm php bin/console doctrine:migrations:migrate -n
```

Load the entity fixtures

```bash
docker exec -it codesnippets-fpm php bin/console doctrine:fixtures:load
```

Run the linting

```bash
docker exec -it codesnippets-fpm php vendor/bin/php-cs-fixer fix
```
