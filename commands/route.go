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
		fmt.Printf("%-8s%s\n", colorizeMethod(r.Method), colorizeRoute(r.Prefix+r.Route))
	}

	return nil
}

func colorizeMethod(method string) string {

	color := "white"
	if method == "GET" {
		color = "blue"
	} else if method == "POST" || method == "PUT" || method == "PATCH" {
		color = "yellow"
	} else if method == "DELETE" {
		color = "red"
	}
	coloredMethod := ui.Output(ui.NewColor(color, "default"), method)
	escapeCodesLength := len(coloredMethod) - len(method)
	formattingWidth := 8 + escapeCodesLength

	return fmt.Sprintf("%-*s", formattingWidth, coloredMethod)
}

func colorizeRoute(route string) string {
	parts := utils.GetPartsByField(route)
	var result string
	for _, part := range parts {
		if part.IsField {
			result = result + ui.Output(ui.NewColor("yellow", "default"), part.Part)
		} else {
			result = result + part.Part
		}
	}
	return result
}
