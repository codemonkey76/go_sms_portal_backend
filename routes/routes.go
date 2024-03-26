package routes

import (
	"net/http"
	"sms_portal/database"
	"sms_portal/db/sqlc"
	"sms_portal/routes/api"
	"sms_portal/utils"
)

func RegisterRoutes(mux *http.ServeMux) *utils.RouteRegistrar {
	db := database.GetDB()
	queries := sqlc.New(db)

	rr := utils.NewRouteRegistrar(mux, utils.WithDB(db, queries))

	api.RegisterRoutes("/api/v1", rr)

	return rr
}
