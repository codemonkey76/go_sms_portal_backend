package main

import (
	"flag"
	"os"
	"sms_portal/commands"
	"sms_portal/internal/database"
)

func main() {
	db := database.GetDB()
	defer db.Close()

	flag.Usage = commands.Usage
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	cmd, args := flag.Args()[0], flag.Args()[1:]
	commands.RunCommand(cmd, args)
}
