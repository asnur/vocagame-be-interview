STEP ?= -all

run:
	go run ./cmd/http/main.go

migrate:
	SQL_COMMAND=$(NAME) docker compose -f docker-compose-migrate.yml run --rm migrate

migrate.up:
	docker compose -f docker-compose-migrate.yml run --rm migrate-up
	
migrate.down:
	MIGRATION_STEP=$(STEP) docker compose -f docker-compose-migrate.yml run --rm migrate-down

# Test commands
test:
	./run_tests.sh

test.unit:
	go test ./internal/usecase/wallet/... -v

test.integration:
	go test ./test/integration/... -v

test.deposit:
	go test ./internal/usecase/wallet/deposit_test.go -v

test.transfer:
	go test ./internal/usecase/wallet/transfer_test.go -v

test.coverage:
	go test ./internal/usecase/wallet/... -cover

# Build commands
build:
	go build -o bin/app ./cmd/http/main.go

clean:
	rm -rf bin/

# Docker commands
docker.build:
	docker build -t vocagame-wallet .

docker.run:
	docker run -p 3030:3030 vocagame-wallet

# Development commands
deps:
	go mod download
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

lint: fmt vet

# Help command
help:
	@echo "Available commands:"
	@echo "  run              - Run the application"
	@echo "  build            - Build the application"
	@echo "  clean            - Clean build artifacts"
	@echo "  test             - Run all tests"
	@echo "  test.unit        - Run unit tests only"
	@echo "  test.integration - Run integration tests only"
	@echo "  test.deposit     - Run deposit tests only"
	@echo "  test.transfer    - Run transfer tests only"
	@echo "  test.coverage    - Run tests with coverage"
	@echo "  migrate.up       - Run database migrations"
	@echo "  migrate.down     - Rollback migrations"
	@echo "  docker.build     - Build Docker image"
	@echo "  docker.run       - Run in Docker"
	@echo "  deps             - Download dependencies"
	@echo "  lint             - Run linters"
	@echo "  help             - Show this help"