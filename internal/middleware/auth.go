package middleware

import (
	"blog/internal/contextutil"
	"blog/internal/service"
	"context"
	"net/http"
)

type Middleware struct {
	Middleware service.AuthService
}

func NewAuthMiddleware(service service.AuthService) *Middleware {
	return &Middleware{Middleware: service}
}

func (md Middleware) SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		next.ServeHTTP(w, r)
	})
}

func (md Middleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		cookie, err := r.Cookie("jwt-token")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := md.Middleware.Validate_Token(cookie.Value)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), contextutil.UserIDKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
