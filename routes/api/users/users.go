package users

import (
	"context"
	"log"
	"net/http"
	"sms_portal/database"
	"sms_portal/db/sqlc"
	"sms_portal/pagination"
	"sms_portal/utils"
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

	users := [][]string{
		{"John Doe", "john@example.com", "password"},
		{"Jane Doe", "jane@example.com", "password"},
		{"John Smith", "jsmith@example.com", "password"},
	}

	for _, user := range users {
		_, err := queries.CreateUser(ctx, sqlc.CreateUserParams{
			Name:     user[0],
			Email:    user[1],
			Password: user[2],
		})

		if err != nil {
			log.Printf("Could not add user: %s", err)
		}
	}
}
