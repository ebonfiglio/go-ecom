package http

import (
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/ebonfiglio/go-ecom/pkg/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter(verifier *oidc.IDTokenVerifier) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/v1")

	api.Use(auth.Middleware(verifier))
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

	return r
}
