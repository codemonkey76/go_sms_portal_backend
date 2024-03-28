package seed

import (
	"context"
	"fmt"
	"sms_portal/db/sqlc"
	"sms_portal/internal/auth"
	"sms_portal/internal/database"
	"sms_portal/internal/ui"

	"github.com/brianvoe/gofakeit"
)

func SeedUsers() {
	db := database.GetDB()
	defer db.Close()

	queries := sqlc.New(db)
	ctx := context.Background()

	for i := 0; i < 5; i++ {
		name := gofakeit.Name()
		email := gofakeit.Email()
		password := string(auth.HashPassword("password"))

		u, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
			Name:     name,
			Email:    email,
			Password: password,
		})

		if err != nil {
			ui.Error(fmt.Sprintf("Could not add user: %s", err))
		} else {
			ui.Info(fmt.Sprintf("Added user: %s - %s", u.Email, password))
		}
	}
}

func SeedAll() {
	SeedUsers()
}
