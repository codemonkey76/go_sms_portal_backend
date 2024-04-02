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

func Cors(w http.ResponseWriter, r *http.Request, deps utils.HandlerDependencies) (interface{}, error) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Content-Type", "application/json")
	return nil, nil

}
func RegisterRoutes(mux *http.ServeMux) *utils.RouteRegistrar {
	db := database.GetDB()
	queries := sqlc.New(db)

	rr := utils.NewRouteRegistrar(mux, utils.WithDB(db, queries))
	baseStack := middleware.CreateStack(middleware.CorsHandler, middleware.LogRequestHandler)
	authStack := middleware.CreateStack(middleware.AuthHandler, middleware.CorsHandler, middleware.LogRequestHandler)
	prefix := "/api/v1"

	rr.AddHandler("POST", prefix, "/auth/login", rpc_auth.AuthLogin, baseStack)
	rr.AddHandler("OPTIONS", prefix, "/auth/login", Cors, baseStack)
	rr.AddHandler("POST", prefix, "/auth/logout", rpc_auth.AuthLogout, authStack)
	rr.AddHandler("GET", prefix, "/users/", rpc_users.UsersIndex, authStack)
	rr.AddHandler("GET", prefix, "/users/{id}", rpc_users.UsersGet, authStack)
	rr.AddHandler("GET", prefix, "/auth/check-session", rpc_auth.CheckSession, authStack)
	rr.AddHandler("OPTIONS", prefix, "/auth/check-session", Cors, authStack)

	return rr
}
