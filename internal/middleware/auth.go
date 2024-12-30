package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("User-ID") == "" {
			http.Error(w, "User-ID header is required", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
