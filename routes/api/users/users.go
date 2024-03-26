package users

import (
	"errors"
	"net/http"
	"sms_portal/database"
	"sms_portal/db/sqlc"
	"sms_portal/routes"
)

func RegisterRoutes(prefix string, rr routes.RouteRegistrar) {
	rr.AddHandler("GET", prefix, "/", Index)

}

func Index(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	defer db.Close()

	w.Write([]byte("Hello, World!"))
	w.WriteHeader(http.StatusOK)
}

func Get(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	db := database.GetDB()
	defer db.Close()
	queries := sqlc.New(db)

	user, err := queries.GetUserById(r.Context(), 1)
	if err != nil {
		return nil, errors.New("Model not found")
	}
	return user, nil
}
