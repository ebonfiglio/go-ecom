package main

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/ebonfiglio/go-ecom/services/bff/internal/http"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	domain := os.Getenv("AUTH0_DOMAIN")
	audience := os.Getenv("AUTH0_AUDIENCE")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	if domain == "" || audience == "" || clientID == "" {
		log.Fatal("AUTH0_DOMAIN, AUTH0_AUDIENCE, AUTH0_CLIENT_ID must be set")
	}

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://"+domain+"/")
	if err != nil {
		log.Fatalf("failed to initialize OIDC provider: %v", err)
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: audience})

	r := http.NewRouter(verifier)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("BFF running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
