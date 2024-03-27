package users

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"sms_portal/auth"
	"sms_portal/database"
	"sms_portal/db/sqlc"
	httperrors "sms_portal/http"
	"sms_portal/pagination"
	"sms_portal/utils"
	"strconv"

	"github.com/brianvoe/gofakeit/v6"
)

func RegisterRoutes(prefix string, rr *utils.RouteRegistrar) {
	rr.AddHandler("GET", prefix, "/", Index, nil)
	rr.AddHandler("GET", prefix, "/{id}", Get, nil)
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
	// Extract path variable
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return nil, httperrors.NotFound()
	}

	// Get User by ID
	user, err := deps.Queries.GetUserById(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, httperrors.NotFound()
		}
		return nil, httperrors.InternalServerError()
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
