package http

import (
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

// NewRouter sets up routes and middleware
func NewRouter(verifier *oidc.IDTokenVerifier) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Health endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "bff ok"})
	})

	api := r.Group("/v1")
	// Apply Auth middleware globally to /v1
	api.Use(AuthMiddleware(verifier))
	{
		// Placeholder proxy endpoints
		api.GET("/users/:id", func(c *gin.Context) {
			// TODO: forward to User Service
			c.JSON(http.StatusOK, gin.H{"user_id": c.Param("id")})
		})
		api.POST("/users", func(c *gin.Context) {
			// TODO: orchestrate signup with Auth0 & user DB
			c.JSON(http.StatusCreated, gin.H{"id": "<new-id>"})
		})
	}

	return r
}
