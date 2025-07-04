# Simple Makefile for a Go project

# Build the application
all: build test

# Run the application
run:
	@go run cmd/api/main.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

lint:
	@echo "Linting..."
	@golangci-lint run ./...

lint-fix:
	@echo "Linting..."
	@golangci-lint run ./... --fix

order:
	@air -c ./tools/air-configs/order.air.toml

promotion:
	@air -c ./tools/air-configs/promotion.air.toml

protoc:
	@echo "Generating protobuf code..."
	@find ./proto -name "*.proto" -exec \
	protoc --proto_path=./proto \
	--go_out=./proto/pb --go_opt=paths=source_relative \
	--go-grpc_out=./proto/pb --go-grpc_opt=paths=source_relative \
	{} \;

up:
	@COMPOSE_BAKE=true docker compose -f docker-compose.dev.yml up -d

down:
	@docker compose -f docker-compose.dev.yml down

build:
	@echo "Building images ..."
	@for service in $(SERVICE); do \
		docker build -f services/$$service/Dockerfile -t fbin243/ztf-$$service:latest .; \
		docker push fbin243/ztf-$$service:latest; \
		echo "Built and pushed fbin243/ztf-$$service:latest"; \
	done

k8s-up:
	@chmod +x ./scripts/zcs_deploy.sh
	@./scripts/zcs_deploy.sh

k8s-down:
	@chmod +x ./scripts/zcs_delete.sh
	@./scripts/zcs_delete.sh

.PHONY: all build run test clean watch lint lint-fix order promotion protoc k8s-up k8s-down