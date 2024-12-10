# load .env
include .env

start:
	docker compose up -d
stop:
	docker compose stop
down:
	docker compose down

migrate-up:
	migrate \
		-path db/migration \
		-database "postgresql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@$(DATABASE_HOST):5432/$(DATABASE_NAME)?sslmode=disable" \
		--verbose up

migrate-down:
	migrate \
		-path db/migration \
		-database "postgresql://$(DATABASE_USERNAME):$(DATABASE_PASSWORD)@$(DATABASE_HOST):5432/$(DATABASE_NAME)?sslmode=disable" \
		--verbose down

sqlc-gen:
	sqlc generate
