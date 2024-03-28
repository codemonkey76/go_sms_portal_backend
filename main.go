package main

import (
	"fmt"
	"os"
	"sms_portal/internal/commands"
	"sms_portal/internal/database"
)

func main() {
	// Check the Database conenction, so we can fail early
	db := database.GetDB()
	defer db.Close()

	if len(os.Args) < 2 {
		fmt.Println("Expected a subcommand")
		commands.HelpCommand()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "genkey":
		commands.GenerateKeyCommand()
	case "route":
		commands.RouteCommand()
	case "serve":
		commands.ServeCommand()
	case "help":
		commands.HelpCommand()
	case "seed":
		commands.SeedCommand()

	}
}
