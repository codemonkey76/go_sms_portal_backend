package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sms_portal/db/sqlc"
	"sms_portal/internal/auth"
	"sms_portal/internal/database"
	"sms_portal/internal/env"
	"sms_portal/internal/ui"
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

func LogRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		ui.Info(wrapped.statusCode, ui.ColorizeUri(r.Method, r.URL.Path), time.Since(start).String())
	})
}

func AuthHandler(next http.Handler) http.Handler {
	// AuthHandler is a middleware that checks if the request has a valid token
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate Token
		email, err := validateToken(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Attach user to request context
		ctx := context.WithValue(r.Context(), "userFunc", UserFunc(func() sqlc.User {
			db := database.GetDB()
			queries := sqlc.New(db)
			user, _ := queries.GetUserByEmail(r.Context(), email)
			return sqlc.User{
				ID:              user.ID,
				Name:            user.Name,
				Email:           user.Email,
				EmailVerifiedAt: user.EmailVerifiedAt,
				Password:        "",
				Active:          user.Active,
				CreatedAt:       user.CreatedAt,
				UpdatedAt:       user.UpdatedAt,
			}
		}))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// region Helper Functions

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func getTokenFromHeader(w http.ResponseWriter, r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header is required")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || strings.ToLower(headerParts[0]) != "bearer" {
		return "", errors.New("Authorization header must be in 'Bearer <token>' format")
	}

	return headerParts[1], nil
}

func validateToken(w http.ResponseWriter, r *http.Request) (string, error) {
	tokenStr, err := getTokenFromHeader(w, r)
	if err != nil {
		return "", err
	}

	claims := &auth.Claims{}

	secret := env.Env(".env").Get("APP_KEY", "")
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", errors.New("Invalid token signature")
		}
		return "", errors.New(fmt.Sprintf("Error parsing token: %s", err.Error()))
	}

	if !token.Valid {
		return "", errors.New("Invalid or expired token")
	}

	ui.Info("Authenticated user: ", claims.Email)
	return claims.Email, nil
}

//endregion

type UserFunc func() sqlc.User
