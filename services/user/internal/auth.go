package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
)

type Middleware struct {
	verifier *oidc.IDTokenVerifier
}

func NewMiddleware(ctx context.Context, domain, audience, clientID string) (*Middleware, error) {
	provider, err := oidc.NewProvider(ctx, "https://"+domain+"/")
	if err != nil {
		return nil, err
	}
	config := &oidc.Config{ClientID: clientID, SkipClientIDCheck: false}
	verifier := provider.Verifier(config)
	return &Middleware{verifier: verifier}, nil
}

func (m *Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(w, "missing auth", http.StatusUnauthorized)
			return
		}
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}
		_, err := m.verifier.Verify(r.Context(), parts[1])
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
