package routes

import (
	"net/http"
	"sms_portal/db/sqlc"
	"sms_portal/internal/database"
	"sms_portal/internal/middleware"
	"sms_portal/internal/rpc_auth"
	"sms_portal/internal/rpc_users"
	"sms_portal/internal/utils"
)

func RegisterRoutes(mux *http.ServeMux) *utils.RouteRegistrar {
	db := database.GetDB()
	queries := sqlc.New(db)

	rr := utils.NewRouteRegistrar(mux, utils.WithDB(db, queries))
	baseStack := middleware.CreateStack(middleware.CorsHandler, middleware.LogRequestHandler)
	authStack := middleware.CreateStack(middleware.AuthHandler, middleware.LogRequestHandler)
	prefix := "/api/v1"

	rr.AddHandler("POST", prefix, "/auth/login", rpc_auth.AuthLogin, baseStack)
	rr.AddHandler("POST", prefix, "/auth/logout", rpc_auth.AuthLogout, authStack)
	rr.AddHandler("GET", prefix, "/users/", rpc_users.UsersIndex, authStack)
	rr.AddHandler("GET", prefix, "/users/{id}", rpc_users.UsersGet, authStack)

	return rr
}
