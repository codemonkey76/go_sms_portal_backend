package main

import (
	"context"
	"flag"
	"os"
	"sms_portal/db/sqlc"
	"sms_portal/internal/commands"
	"sms_portal/internal/config"
	"sms_portal/internal/database"
	"sms_portal/internal/ui"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func main() {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	j, err := s.NewJob(
		gocron.DurationJob(24*time.Hour),
		gocron.NewTask(func() {
			db := database.GetDB()
			queries := sqlc.New(db)
			queries.DeleteExpiredSessions(context.Background(), time.Now().Unix()-config.SessionExpiration*60)
		}),
	)
	if err != nil {
		panic(err)
	}

	ui.Info("Scheduled Job: ", "ClearExpiredSessions", j.ID())
	s.Start()
	defer s.Shutdown()

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
