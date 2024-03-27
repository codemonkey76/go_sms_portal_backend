package middleware

import (
	"fmt"
	"net/http"
	"sms_portal/env"
	"sms_portal/routes/api/auth"
	"sms_portal/ui"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(xs) - 1; i >= 0; i-- {
			x := xs[i]
			next = x(next)
		}
		return next
	}
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LogRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		ui.Info(wrapped.statusCode, ui.ColorizeUri(r.Method, r.URL.Path), time.Since(start).String())
	})
}

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ui.Info("AuthHandler")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Split the header into "Bearer" and the <token>
		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
			http.Error(w, "Authorization header must be in 'Bearer <token>' format", http.StatusUnauthorized)
			return
		}

		tokenStr := headerParts[1]

		// Here, you should add your token validation logic. For now, let's just log the token.
		// Log the token or proceed with further token validation.

		// If token validation passes, call the next handler; otherwise, return an error.
		// This is a placeholder for your token validation logic.
		claims := &auth.Claims{}

		secret := env.Env(".env").Get("APP_KEY", "")
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, fmt.Sprintf("Error parsing token: %s", err.Error()), http.StatusBadRequest)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}
		fmt.Println("Authenticated user:", claims.Email)

		// Do some authentication stuff here
		next.ServeHTTP(w, r)
	})
}
