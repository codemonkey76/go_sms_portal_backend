package seed

import (
	"context"
	"fmt"
	"sms_portal/db/sqlc"
	"sms_portal/internal/auth"
	"sms_portal/internal/database"
	"sms_portal/internal/ui"
)

func SeedPermissions() {
	db := database.GetDB()

	queries := sqlc.New(db)
	ctx := context.Background()

	permissions := []string{"users.list", "users.update", "users.store", "users.destroy", "users.get"}

	for _, permission := range permissions {
		p, err := queries.CreatePermission(ctx, permission)

		if err != nil {
			ui.Error(fmt.Sprintf("Could not add permission: %s", err))
		} else {
			ui.Info(fmt.Sprintf("Added permission: %s", p.Name))
		}
	}
}

func SeedUsers() {
	db := database.GetDB()

	queries := sqlc.New(db)
	ctx := context.Background()

	addUser(&ctx, queries, "Admin User", "admin@example.com", "password")
	addUser(&ctx, queries, "Regular User", "user@example.com", "password")
}

func addUser(ctx *context.Context, queries *sqlc.Queries, name, email, password string) {
	hashedPassword := string(auth.HashPassword(password))
	u, err := queries.CreateUser(*ctx, sqlc.CreateUserParams{Name: name, Email: email, Password: hashedPassword})

	if err != nil {
		ui.Error(fmt.Sprintf("Could not add user: %s", err))
	} else {
		ui.Info(fmt.Sprintf("Added user: %s - %s", u.Email, password))
	}
}

func SeedUserPermissions() {
	err := auth.GivePermissionToUserByEmail("admin@example.com", "users.list")
	if err != nil {
		ui.Error(fmt.Sprintf("Could not add permission to user: %s", err))
	} else {
		ui.Info(fmt.Sprintf("Added permission \"%s\" to user \"%s\"", "users.list", "admin@example.com"))
	}
}

func SeedAll() {
	SeedUsers()
	SeedPermissions()
	SeedUserPermissions()
}
