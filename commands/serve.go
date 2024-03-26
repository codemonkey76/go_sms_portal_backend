package commands

import (
	"log"
	"net/http"
	"sms_portal/routes"
	"sms_portal/ui"
)

func ServeCommand(args []string) error {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	log.Println(ui.Colorize(ui.ColorGreen, "Starting server on port 8080"))
	http.ListenAndServe(":8080", mux)

	return nil
}
