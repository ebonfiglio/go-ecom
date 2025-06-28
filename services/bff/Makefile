APP_NAME := go-ecom-bff
CMD_PATH := cmd/bff
ENV_FILE := .env

.PHONY: build run test fmt lint docker clean

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

run: build
	@set -a; . $(ENV_FILE); set +a;
	./bin/$(APP_NAME)

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./...

docker:
	docker build -t $(APP_NAME):latest .

clean:
	go clean
	rm -rf bin/$(APP_NAME)