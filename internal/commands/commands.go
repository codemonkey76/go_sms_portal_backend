package commands

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sms_portal/db/sqlc"
	"sms_portal/internal/config"
	"sms_portal/internal/database"
	"sms_portal/internal/env"
	"sms_portal/internal/middleware"
	"sms_portal/internal/routes"
	"sms_portal/internal/seed"
	"sms_portal/internal/ui"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-co-op/gocron/v2"
)

type Command struct {
	Name string
	Help string
}

var commands = []Command{
	{Name: "genkey", Help: "Generate a new key"},
	{Name: "seed", Help: "Seed the database with some data"},
	{Name: "route", Help: "List all the api routes"},
	{Name: "serve", Help: "Start the web server"},
	{Name: "help", Help: "Print this help message"},
}

func setupJobScheduler() (gocron.Scheduler, error) {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	_, err = s.NewJob(
		gocron.DurationJob(24*time.Hour),
		gocron.NewTask(func() {
			db := database.GetDB()
			queries := sqlc.New(db)
			queries.DeleteExpiredSessions(context.Background(), time.Now().Unix()-config.SessionExpiration*60)
		}),
	)
	if err != nil {
		return nil, err
	}

	ui.Info("Scheduled Job: ", "ClearExpiredSessions")
	s.Start()

	return s, nil
}

func ServeCommand() {
	var mu sync.Mutex
	currentPort := env.AppConfig.ServerPort
	ui.Info("Setting currentPort to: ", currentPort)
	serverDone := make(chan bool)
	ui.Info("Launching Serve go routine.")
	go Serve(currentPort, serverDone)
	ui.Info("Server should now be up and running.")

	ui.Info("Setting up file watcher.")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		ui.Error("Error creating watcher: ", err.Error())
	}
	defer watcher.Close()

	debounceTimer := time.NewTimer(1 * time.Second)
	debounceTimer.Stop()

	go func() {
		for {
			ui.Info("Checking for watcher events.")
			select {
			case event, ok := <-watcher.Events:
				ui.Info("Received event: ", event)
				if !ok {
					return
				}
				if event.Op&(fsnotify.Write|fsnotify.Rename) != 0 {
					debounceTimer.Reset(500 * time.Millisecond)
				}
			case <-debounceTimer.C:
				mu.Lock()
				defer mu.Unlock()
				env.Init()

				if env.AppConfig.ServerPort != currentPort {
					ui.Info(".env file modified, SERVER_PORT Changed. Restarting server.")
					currentPort = env.AppConfig.ServerPort

					ui.Info("Sending done signal.")
					serverDone <- true
					<-serverDone
					serverDone = make(chan bool)
					go Serve(currentPort, serverDone)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				ui.Error("Error watching file: ", err.Error())
			}
		}
	}()

	err = watcher.Add(".env")
	if err != nil {
		ui.Error("Error watching file: ", err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ui.Info("Received quit signal. Sending done signal.")
	serverDone <- true

}

func Serve(port string, done chan bool) {
	ui.Info("Setting up job scheduler.")
	s, err := setupJobScheduler()
	if err != nil {
		ui.Error(fmt.Sprintf("Error setting up job scheduler: %s", err.Error()))
		defer s.Shutdown()
	}

	ui.Info("Registering Application routes.")
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)
	server := http.Server{
		Addr:    ":" + port,
		Handler: middleware.LogRequestHandler(mux),
	}

	ui.Info("Server started on port " + port)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			ui.Error("Fatal error occurred running server: ", err.Error())
		}
	}()
	ui.Info("Go routine started. waiting on done channel.")
	<-done
	ui.Info("Received done signal. Shutting down server.")

	// Create a context with countdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		ui.Error("Error shutting down server: ", err.Error())
	} else {
		ui.Info("Server stopped.")
	}
	done <- true
}

func SeedCommand() {
	env.Init()
	seedCmd := flag.NewFlagSet("seed", flag.ExitOnError)
	var entities string
	seedCmd.StringVar(&entities, "entities", "", "Comma-separated list of entities to seed: all,users,tasks,customers,permissions,roles")

	if len(os.Args) > 2 {
		seedCmd.Parse(os.Args[2:])
	} else {
		ui.Error("The 'seed' command requires the -entities flag to be set.")
		seedCmd.PrintDefaults()
		os.Exit(1)
	}

	entitiesList := strings.Split(entities, ",")
	for _, entity := range entitiesList {
		switch entity {
		case "all":
			seed.SeedAll()
		case "users":
			seed.SeedUsers()
		case "permissions":
			seed.SeedPermissions()
		default:
			ui.Warn("Invalid entity: ", entity)
		}
	}

}

func RouteCommand() {
	mux := http.NewServeMux()
	rr := routes.RegisterRoutes(mux)

	for _, r := range rr.Routes {
		fmt.Println(ui.ColorizeUri(r.Method, r.Prefix+r.Route))
	}
}

func HelpCommand() {
	green := ui.NewColor("green", "default", "bold")
	yellow := ui.NewColor("yellow", "default", "bold")

	fmt.Print(ui.Output(green, os.Args[0]))
	fmt.Print(" version ")
	fmt.Print(ui.Output(yellow, config.ReleaseVersion))
	fmt.Printf(" %s\n\n", config.ReleaseDate)
	fmt.Print(ui.Output(yellow, "Usage:\n"))
	fmt.Print("    command [subcommand]\n\n")
	fmt.Print(ui.Output(yellow, "Available Sub-Commands:\n"))
	for _, c := range commands {
		name := ui.Output(green, c.Name)
		escapeCodeLen := len(name) - len(c.Name)
		formatWidth := 8 + escapeCodeLen
		fmt.Printf("    %-*s", formatWidth, name)
		fmt.Println(c.Help)
	}
}

func GenerateKeyCommand() {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	env.Env(".env").Set("APP_KEY", string(b))
	ui.Info("Key generated successfully.")
}
