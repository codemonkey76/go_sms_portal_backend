package seed

import (
	"context"
	"log"
	"sms_portal/db/sqlc"
	"sms_portal/internal/auth"
	"sms_portal/internal/database"

	"github.com/brianvoe/gofakeit"
)

func SeedUsers() {
	db := database.GetDB()
	defer db.Close()

	queries := sqlc.New(db)
	ctx := context.Background()

	for i := 0; i < 100; i++ {
		name := gofakeit.Name()
		email := gofakeit.Email()
		password := string(auth.HashPassword("password"))

		u, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
			Name:     name,
			Email:    email,
			Password: password,
		})

		if err != nil {
			log.Printf("Could not add user: %s", err)
		} else {
			log.Printf("Added user: %s - %s", u.Email, password)
		}
	}
}
