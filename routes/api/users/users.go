package users

import (
	"context"
	"log"
	"net/http"
	"sms_portal/auth"
	"sms_portal/database"
	"sms_portal/db/sqlc"
	"sms_portal/pagination"
	"sms_portal/utils"

	"github.com/brianvoe/gofakeit/v6"
)

func RegisterRoutes(prefix string, rr *utils.RouteRegistrar) {
	rr.AddHandler("GET", prefix, "/", Index)
	rr.AddHandler("GET", prefix, "/{id}", Get)
}

func Index(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	page, perPage, search := pagination.GetPaginationOptions(r)

	users, err := deps.Queries.ListUsers(r.Context(), sqlc.ListUsersParams{
		Search: search,
		Offset: int32((page - 1) * perPage),
		Limit:  int32(perPage),
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func Get(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	user, err := deps.Queries.GetUserById(r.Context(), 1)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Seed() {
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
