package http

import (
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks JWT and aborts if invalid
func AuthMiddleware(verifier *oidc.IDTokenVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		head := c.GetHeader("Authorization")
		if len(head) < 7 || head[:7] != "Bearer " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or malformed auth header"})
			return
		}
		tok := head[7:]
		if _, err := verifier.Verify(c.Request.Context(), tok); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}
