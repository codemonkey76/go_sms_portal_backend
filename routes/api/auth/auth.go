package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	rr.AddHandler("POST", prefix, "/login", Login)
	rr.AddHandler("POST", prefix, "/logout", Logout)
}

func Login(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		return nil, err
	}

	user, err := deps.Queries.GetUserByEmail(r.Context(), creds.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return nil, err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Printf("%s", token.Raw)
	return nil, nil

}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func Logout(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	return nil, nil
}
