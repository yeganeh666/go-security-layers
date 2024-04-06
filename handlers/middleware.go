package handlers

import (
	"fmt"
	"net/http"
	"security-example/pkg/auth"
	"security-example/pkg/log"
	"security-example/pkg/ratelimiter"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		username, err := auth.ValidateToken(token)
		if err != nil {
			http.Error(w, fmt.Sprint("Unauthorized ", err), http.StatusUnauthorized)
			return
		}

		if !auth.IsAuthorized(username, r.URL.Path) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if !ratelimiter.Check(r.RemoteAddr) {
			http.Error(w, "Rate Limit Exceeded", http.StatusTooManyRequests)
			return
		}

		log.AuditEvent(r.Method, r.URL.Path, username, r.RemoteAddr)

		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		//w.Header().Set("Content-Security-Policy", "script-src 'self'; img-src 'self' https://example.com")

		r = r.WithContext(r.Context())
		next(w, r)
	}
}
