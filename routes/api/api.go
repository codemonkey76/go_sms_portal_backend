package api

import (
	"sms_portal/routes"
	"sms_portal/routes/api/users"
)

func RegisterRoutes(prefix string, rr *routes.RouteRegistrar) {
	users.RegisterRoutes(prefix+"/users", rr)
}
