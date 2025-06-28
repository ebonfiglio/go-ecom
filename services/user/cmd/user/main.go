package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	authMiddleware := func(c *gin.Context) {
		hdr := c.GetHeader("Authorization")
		parts := strings.SplitN(hdr, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
			return
		}
		if _, err := verifier.Verify(c.Request.Context(), parts[1]); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}

	api := r.Group("/v1")
	api.Use(authMiddleware)
	{
		api.POST("/echo", func(c *gin.Context) {
			var payload struct {
				Message string `json:"message" binding:"required"`
			}
			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"echo": payload.Message})
		})

	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
