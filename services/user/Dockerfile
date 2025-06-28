FROM golang:1.24.2-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w -X main.version=${VERSION}" \
    -o /app/bin/go-ecom-user ./services/user/cmd/user/main.go

FROM scratch
WORKDIR /app

COPY --from=builder /app/bin/go-ecom-user .

USER 65532:65532

EXPOSE 8080

ENTRYPOINT ["/app/go-ecom-user"]
