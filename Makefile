#!make
include local.env

up-stack:
	docker-compose --env-file local.env up --build -d

up:
	docker-compose --env-file local.env up -d

down:
	docker-compose --env-file local.env down

create-migrate:
	tern new ${name} -m migrate/data

.PHONY: up-stack up down create-migrate