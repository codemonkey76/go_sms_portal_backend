package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sms_portal/db/sqlc"
	"sms_portal/internal/config"
	"sms_portal/internal/database"
	"sms_portal/internal/env"
	"sms_portal/internal/ui"
	"time"
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

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ui.Info("CORS Middleware enabled")

		w.Header().Set("Access-Control-Allow-Origin", env.AppConfig.FrontendDomain)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func LogRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ui.Info("LogRequest Middleware enabled")

		start := time.Now()

		wrapped := &wrappedWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(wrapped, r)
		uri := ui.ColorizeUri(r.Method, r.URL.Path)

		s := fmt.Sprintf("%d %s %s", wrapped.statusCode, uri, time.Since(start).String())
		if wrapped.statusCode == http.StatusOK {
			ui.Info(s)
		} else {
			ui.Error(s)
		}
	})
}

func AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ui.Info("Auth Middleware enabled")

		user_id, err := validateSessionToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", user_id)

		next.ServeHTTP(w, r.WithContext(ctx))
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

func validateSessionToken(r *http.Request) (int64, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return -1, errors.New("Session token not found")
	}

	sessionToken := cookie.Value

	db := database.GetDB()
	queries := sqlc.New(db)

	session, err := queries.GetSessionByToken(r.Context(), sessionToken)
	if err != nil {
		return -1, errors.New("Invalid session token")
	}

	// Last Activity stores the time the session was last used by using Unix time
	// Expiry is set by the const config.SessionExpiration in minutes

	if time.Now().Unix()-session.LastActivity > config.SessionExpiration*60 {
		return -1, errors.New("Session token has expired")
	}

	return session.UserID, nil

}

//endregion
