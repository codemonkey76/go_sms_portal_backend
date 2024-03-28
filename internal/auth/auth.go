package auth

import (
	"context"
	"sms_portal/db/sqlc"
	"sms_portal/internal/database"

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

func HasPermission(id int64, permission string) bool {
	db := database.GetDB()
	queries := sqlc.New(db)

	// Check if user has directly assigned permission
	permissions, err := queries.ListUserPermissions(context.Background(), id)
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
