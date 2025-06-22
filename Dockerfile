# Start from the official Go image
FROM golang:1.24.2-alpine AS builder

# Set the working directory
WORKDIR /go-ecom

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o go-ecom ./cmd/main.go

# Use a minimal final image
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder
COPY --from=builder /go-ecom/go-ecom .

# Command to run the app
CMD ["./go-ecom"]
