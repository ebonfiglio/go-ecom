package main

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/ebonfiglio/go-ecom/services/user/internal/http"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on environment")
	}
}

func main() {
	domain := os.Getenv("AUTH0_DOMAIN")
	audience := os.Getenv("AUTH0_AUDIENCE")
	clientID := os.Getenv("AUTH0_CLIENT_ID")
	if domain == "" || audience == "" || clientID == "" {
		log.Fatal("AUTH0_DOMAIN, AUTH0_AUDIENCE, AUTH0_CLIENT_ID must be set")
	}

	oidcCtx := context.Background()
	provider, err := oidc.NewProvider(oidcCtx, "https://"+domain+"/")
	if err != nil {
		log.Fatalf("failed to init OIDC provider: %v", err)
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: audience})

	r := http.NewRouter(verifier)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)

	log.Printf("User Service running on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
