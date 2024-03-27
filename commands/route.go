package commands

import (
	"fmt"
	"net/http"
	"sms_portal/routes/api"
	"sms_portal/ui"
	"sms_portal/utils"
)

func RouteCommand(args []string) error {
	mux := http.NewServeMux()
	rr := utils.NewRouteRegistrar(mux)
	api.RegisterRoutes("/api/v1", rr)

	for _, r := range rr.Routes {
		fmt.Printf("%-8s%s\n", ui.ColorizeMethod(r.Method), ui.ColorizeRoute(r.Prefix+r.Route))
	}

	return nil
}
