FROM golang:1.24.2-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w -X main.version=${VERSION}" -o bin/go-ecom-bff cmd/bff

FROM scratch
WORKDIR /app
COPY --from=builder /app/bin/go-ecom-bff .
ENV PORT=8081
EXPOSE 8081
USER 65532:65532
ENTRYPOINT ["/app/go-ecom-bff"]
