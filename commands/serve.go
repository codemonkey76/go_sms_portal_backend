package commands

import (
	"fmt"
	"net/http"
	"sms_portal/env"
	"sms_portal/http/middleware"
	"sms_portal/routes"
	"sms_portal/ui"
)

func ServeCommand(args []string) error {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	port := env.Env(".env").Get("SERVER_PORT", "8080")
	server := http.Server{
		Addr:    ":" + port,
		Handler: middleware.LogRequestHandler(mux),
	}
	ui.Info("Server started on port " + port)
	err := server.ListenAndServe()
	ui.Error(fmt.Sprintf("Error occurred while starting server: %s", err.Error()))
	return nil
}
