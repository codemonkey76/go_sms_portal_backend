package commands

import (
	"fmt"
	"net/http"
	"sms_portal/routes"
	"sms_portal/ui"
)

func ServeCommand(args []string) error {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	ui.Info("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	ui.Error(fmt.Sprintf("Error occurred while starting server: %s", err.Error()))
	return nil
}
