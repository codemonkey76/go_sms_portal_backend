package api

import (
	"sms_portal/routes/api/users"
	"sms_portal/utils"
)

func RegisterRoutes(prefix string, rr *utils.RouteRegistrar) {
	users.RegisterRoutes(prefix+"/users", rr)
}
