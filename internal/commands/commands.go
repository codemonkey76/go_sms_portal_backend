package commands

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"slices"
	"sms_portal/internal/config"
	"sms_portal/internal/env"
	"sms_portal/internal/middleware"
	"sms_portal/internal/routes"
	"sms_portal/internal/seed"
	"sms_portal/internal/ui"
	"sms_portal/internal/utils"
)

type Command struct {
	Name string
	Help string
	Run  func(args []string) error
}

var commands = []Command{
	{Name: "genkey", Help: "Generate a new key", Run: GenerateKeyCommand},
	{Name: "seed", Help: "Seed the database with some data", Run: SeedCommand},
	{Name: "route", Help: "List all the api routes", Run: RouteCommand},
	{Name: "serve", Help: "Start the web server", Run: ServeCommand},
	{Name: "help", Help: "Print this help message", Run: HelpCommand},
}

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

func SeedCommand(args []string) error {
	seed.SeedUsers()

	return nil
}

func RouteCommand(args []string) error {
	mux := http.NewServeMux()
	rr := utils.NewRouteRegistrar(mux)
	routes.RegisterRoutes(mux)

	for _, r := range rr.Routes {
		fmt.Printf("%-8s%s\n", ui.ColorizeMethod(r.Method), ui.ColorizeRoute(r.Prefix+r.Route))
	}

	return nil
}

func HelpCommand(arg []string) error {
	return nil
}

func GenerateKeyCommand(args []string) error {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	env.Env(".env").Set("APP_KEY", string(b))
	ui.Info("Key generated successfully.")
	return nil
}

func RunCommand(name string, args []string) {
	cmdIdx := slices.IndexFunc(commands, func(c Command) bool {
		return c.Name == name
	})

	if cmdIdx < 0 {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", name)
		flag.Usage()
		os.Exit(1)
	}

	if err := commands[cmdIdx].Run(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error running command %s: %s\n", name, err)
		os.Exit(1)
	}
}

func Usage() {
	ui.Output(ui.NewColor("green", "black"), os.Args[0])
	ui.Output(ui.NewColor("green", "black"), os.Args[0])
	fmt.Print(" version ")
	ui.Output(ui.NewColor("yellow", "black"), config.ReleaseVersion)
	fmt.Printf(" %s\n\n", config.ReleaseDate)
	ui.Output(ui.NewColor("yellow", "black"), "Usage:\n")
	fmt.Print("    command [options] [arguments]\n\n")
	ui.Output(ui.NewColor("yellow", "black"), "Options:\n")
	ui.Output(ui.NewColor("green", "black"), "    -h, --help                    ")
	fmt.Print("Display help for the given command. When no command is given display help for the list command\n")
	fmt.Print("\n")
	ui.Output(ui.NewColor("yellow", "black"), "Available Commands:\n")
	ui.Output(ui.NewColor("green", "black"), "    seed    ")
	fmt.Print("Seed the database with some data\n")
	ui.Output(ui.NewColor("green", "black"), "    serve   ")
	fmt.Print("Start the server\n")
}
