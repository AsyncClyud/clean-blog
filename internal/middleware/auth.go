package middleware

import (
	"blog/internal/contextutil"
	"blog/internal/service"
	"context"
	"net/http"
)

type AuthMiddleware struct {
	Middleware service.AuthService
}

func NewAuthMiddleware(service service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{Middleware: service}
}

func (md AuthMiddleware) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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
