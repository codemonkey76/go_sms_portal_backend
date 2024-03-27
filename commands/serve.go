package commands

import (
	"fmt"
	"net/http"
	"sms_portal/env"
	"sms_portal/routes"
	"sms_portal/ui"
)

func ServeCommand(args []string) error {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	port := env.Env(".env").Get("SERVER_PORT", "8080")

	ui.Info("Starting server on port: " + port)
	err := http.ListenAndServe(":"+port, mux)
	ui.Error(fmt.Sprintf("Error occurred while starting server: %s", err.Error()))
	return nil
}
