# ─── builder stage ───────────────────────────────────────────────────────────────
FROM golang:1.24.2-alpine AS builder
WORKDIR /app

# Copy and download dependencies only
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your code and build the static binary
COPY . .
ARG VERSION
RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w -X main.version=${VERSION}" \
    -o /app/bin/go-ecom-user ./services/user/cmd/user/main.go

# ─── final stage ────────────────────────────────────────────────────────────────
FROM scratch
WORKDIR /app

# Copy only the binary (no certs or packages)
COPY --from=builder /app/bin/go-ecom-user .

# Drop privileges
USER 65532:65532

# Expose your app port
EXPOSE 8080

ENTRYPOINT ["/app/go-ecom-user"]
