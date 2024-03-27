package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	authtypes "sms_portal/auth"
	"sms_portal/database"
	"sms_portal/db/sqlc"
	"sms_portal/env"
	"sms_portal/ui"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

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

	claims := &authtypes.Claims{}

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

	ui.Info(fmt.Sprintf("Authenticated user: ", claims.Email))
	return claims.Email, nil
}

type UserFunc func() interface{}

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate Token
		email, err := validateToken(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Attach user to request context
		ctx := context.WithValue(r.Context(), "userFunc", UserFunc(func() interface{} {
			db := database.GetDB()
			queries := sqlc.New(db)
			user, _ := queries.GetUserByEmail(r.Context(), email)
			return user
		}))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
