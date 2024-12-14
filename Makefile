# Ensure .env exists
ifneq (,$(wildcard .env))
    include .env
    export
endif

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

test-cover:
	go test ./... -cover -v

test-v:
	go test ./... -v

test-profile:
	go test -coverprofile=coverage.out ./...

test-cover-html:
	go tool cover -html=coverage.out -o coverage.html

open-cover:
	#for macOS
	open coverage.html

open-cover-w:
	$(MAKE) test-profile;
	$(MAKE) test-cover-html;
	$(MAKE) open-cover;

server:
	go run main.go
