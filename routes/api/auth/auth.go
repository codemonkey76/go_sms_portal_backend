package auth

import (
	"encoding/json"
	"net/http"
	"sms_portal/db/sqlc"
	"sms_portal/env"
	httperrors "sms_portal/http"
	"sms_portal/http/middleware"
	"sms_portal/ui"
	"sms_portal/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterRoutes(prefix string, rr *utils.RouteRegistrar) {
	authStack := middleware.CreateStack(middleware.AuthHandler, middleware.LogRequestHandler)
	rr.AddHandler("POST", prefix, "/login", Login, nil)
	rr.AddHandler("POST", prefix, "/logout", Logout, authStack)
}

func Login(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		return nil, httperrors.InvalidCredentials()
	}

	user, err := deps.Queries.GetUserByEmail(r.Context(), creds.Email)
	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	}

	if err != nil {
		return nil, httperrors.InvalidCredentials()
	}
	ui.Info("User logged in: " + creds.Email)

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := env.Env(".env").Get("APP_KEY", "")
	tokenString, err := token.SignedString([]byte(secret))
	return_user := sqlc.GetUserByIdRow{
		ID:              user.ID,
		Name:            user.Name,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
	response := LoginResponse{
		Success:   true,
		Message:   "Login successful.",
		Data:      return_user,
		Token:     tokenString,
		ExpiresIn: expirationTime.Unix(),
	}

	return response, nil
}

type LoginResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Token     string      `json:"token"`
	ExpiresIn int64       `json:"expires_in"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Logout(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	return nil, nil
}
