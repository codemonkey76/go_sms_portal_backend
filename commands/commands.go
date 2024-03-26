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
	{Name: "seed", Help: "Seed the database with some data", Run: SeedCommand},
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
	ui.Colorize(ui.ColorGreen, os.Args[0])
	fmt.Print(" version ")
	ui.Colorize(ui.ColorYellow, config.ReleaseVersion)
	fmt.Printf(" %s\n\n", config.ReleaseDate)
	ui.Colorize(ui.ColorYellow, "Usage:\n")
	fmt.Print("    command [options] [arguments]\n\n")
	ui.Colorize(ui.ColorYellow, "Options:\n")
	ui.Colorize(ui.ColorGreen, "    -h, --help                    ")
	fmt.Print("Display help for the given command. When no command is given display help for the list command\n")
	fmt.Print("\n")
	ui.Colorize(ui.ColorYellow, "Available Commands:\n")
	ui.Colorize(ui.ColorGreen, "    seed    ")
	fmt.Print("Seed the database with some data\n")
	ui.Colorize(ui.ColorGreen, "    serve   ")
	fmt.Print("Start the server\n")
}
