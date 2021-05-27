package main

import (
	"log"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var err error

	cfg, err = config.Load()
	if err != nil {
		return err
	}

	logger = cfg.Logger.MustBuildLogR()

	app := cli.NewApp()
	app.Name = "service"
	app.Commands = []cli.Command{
		{
			Name:   "server",
			Usage:  "start grpc/http server",
			Action: serverAction,
		},
		{
			Name:   "auto-migrate",
			Usage:  "auto-migration database schema",
			Action: autoMigration,
		},
	}

	if app.Run(args) != nil {
		panic(err)
	}
	return nil
}