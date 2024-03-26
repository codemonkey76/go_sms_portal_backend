package commands

import (
	"flag"
	"fmt"
	"os"
	"slices"
	"sms_portal/config"
	"sms_portal/ui"
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
