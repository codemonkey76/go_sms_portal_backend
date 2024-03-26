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
		fmt.Printf("%s%s\n", colorizeMethod(r.Method), colorizeRoute(r.Prefix+r.Route))
	}

	return nil
}

func colorizeMethod(method string) string {
	color := ui.ColorReset
	if method == "GET" {
		color = ui.ColorBlue
	} else if method == "POST" || method == "PUT" || method == "PATCH" {
		color = ui.ColorYellow
	} else if method == "DELETE" {
		color = ui.ColorRed
	}
	return fmt.Sprintf("%s%-8s%s", color, method, ui.ColorReset)
}

func colorizeRoute(route string) string {
	parts := utils.GetPartsByField(route)
	var result string
	for _, part := range parts {
		if part.IsField {
			result = result + fmt.Sprintf("%s%s%s", ui.ColorYellow, part.Part, ui.ColorReset)
		} else {
			result = result + part.Part
		}
	}
	return result
}
