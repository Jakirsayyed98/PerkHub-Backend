include .env.development

# Export all variables from .env
export $(shell sed 's/=.*//' .env.development)

run:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bankapi

# $(MAKE) build-upload

build-upload:
	sudo scp -i ~/F:\PerkHub_Go\perkhub.pem ./perkhub ubuntu@ec2-13-200-72-93.ap-south-1.compute.amazonaws.com:~/perkhub

upload-migration:
	sudo scp -i ~/Desktop/Workspace/paydoh/aws-ec2/paydoh-key.pem -r ./migrations ubuntu@ec2-13-200-72-93.ap-south-1.compute.amazonaws.com:~/bankapi/

git-submodule-update:
	git submodule update --recursive

swagger-generate:
	swag init

# goose migrations
goose-create-new:
	goose -dir ./migrations create $(MIGRATION_NAME) sql
	

create_migration:
	@read -p "Enter migration name: " name; \
	goose -dir ./migrations create $$name sql

goose-up:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://$(POSTGRES_USER_NAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable" goose -dir='./migrations' up

goose-down:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://$(POSTGRES_USER_NAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable" goose -dir='./migrations' down

goose-version:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://$(POSTGRES_USER_NAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/$(POSTGRES_DATABASE)?sslmode=disable" goose -dir='./migrations' version

test_run:
	go test -v -tags unit ./test/unittest
	
test-integration:
	go test -tags integration ./tests/integration_test


test_coverage:
	@go test -cover ./...
	@go tool cover -html=coverage.out

lint:
	golangci-lint run