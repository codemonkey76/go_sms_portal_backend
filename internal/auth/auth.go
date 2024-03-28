package auth

import (
	"context"
	"fmt"
	"sms_portal/db/sqlc"
	"sms_portal/internal/database"
	"sms_portal/internal/ui"

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
	ui.Info(fmt.Sprintf("Checking User %d for permission %s", id, permission))
	db := database.GetDB()
	queries := sqlc.New(db)

	// Check if user has directly assigned permission
	permissions, err := queries.ListUserPermissions(context.Background(), id)
	ui.Info("User has permissions: ", permissions)
	if err != nil {
		return false
	}

	return contains(permissions, permission)
}

func GivePermissionToUserByEmail(email string, permission string) error {
	db := database.GetDB()
	queries := sqlc.New(db)

	u, err := queries.GetUserByEmail(context.Background(), email)
	if err != nil {
		return err
	}

	p, err := queries.GetPermissionByName(context.Background(), permission)
	if err != nil {
		return err
	}

	_, err = queries.AttachPermissionToUser(context.Background(), sqlc.AttachPermissionToUserParams{
		PermissionID: p.ID,
		UserID:       u.ID,
	})
	if err != nil {
		return err
	}

	return nil
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
