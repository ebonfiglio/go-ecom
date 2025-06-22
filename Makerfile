APP_NAME := go-ecom
ENTRY_POINT := ./cmd/main.go

.PHONY: build run test fmt lint docker clean

build:
	go build -o $(APP_NAME) $(ENTRY_POINT)

run:
	go run $(ENTRY_POINT)

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	golangci-lint run

docker:
	docker build -t $(APP_NAME):latest .

clean:
	go clean
	rm -f $(APP_NAME)
