package api

import (
	"sms_portal/routes/api/auth"
	"sms_portal/routes/api/users"
	"sms_portal/utils"
)

func RegisterRoutes(prefix string, rr *utils.RouteRegistrar) {
	auth.RegisterRoutes(prefix+"/auth", rr)
	users.RegisterRoutes(prefix+"/users", rr)
}
