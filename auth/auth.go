package auth

import (
	"context"
	"sms_portal/database"
	"sms_portal/db/sqlc"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plaintextPassword string) []byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func HasPermission(user sqlc.User, permission string) bool {
	db := database.GetDB()
	queries := sqlc.New(db)

	// Check if user has directly assigned permission
	permissions, err := queries.ListUserPermissions(context.Background(), user.ID)
	if err != nil {
		return false
	}

	return contains(permissions, permission)
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
