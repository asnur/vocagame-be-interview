STEP ?= -all

run:
	go run ./cmd/http/main.go

migrate:
	SQL_COMMAND=$(NAME) docker compose -f docker-compose-migrate.yml run --rm migrate

migrate.up:
	docker compose -f docker-compose-migrate.yml run --rm migrate-up
	
migrate.down:
	MIGRATION_STEP=$(STEP) docker compose -f docker-compose-migrate.yml run --rm migrate-down