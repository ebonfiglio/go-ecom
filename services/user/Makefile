APP_NAME := go-ecom-user
ENTRY_POINT := ./services/user/cmd/user/main.go

.PHONY: build run test fmt lint docker clean

VERSION := $(shell git describe --tags --always --dirty)
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"
build:
    go build $(LDFLAGS) -o bin/$(APP_NAME) $(ENTRY_POINT)


run:
	go run $(ENTRY_POINT)

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

ci: fmt lint test
    @echo "🚀 All checks passed!"

docker: build
    docker build \
	      --build-arg VERSION=$(VERSION) \
		     -t $(APP_NAME):$(VERSION) \
			 -t $(APP_NAME):latest \
			.

clean:
	go clean
	rm -f $(APP_NAME)
