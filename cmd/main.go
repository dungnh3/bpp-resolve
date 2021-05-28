package main

import (
	"github.com/dungnh3/bpp-resolve/config"
	"github.com/dungnh3/bpp-resolve/internal/server"
	"github.com/dungnh3/bpp-resolve/internal/service"
	"github.com/go-logr/logr"
	"github.com/urfave/cli"
	"log"
	"os"

	migrateV4 "github.com/golang-migrate/migrate/v4"
	// import mysql
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	// import file
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// import go_bin_data
	_ "github.com/golang-migrate/migrate/v4/source/go_bindata"
)

var (
	cfg    *config.Config
	logger logr.Logger
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
			Name:   "migrate-up",
			Usage:  "auto migration up",
			Action: migrateUp,
		},
		{
			Name:   "migrate-down",
			Usage:  "auto migration down",
			Action: migrateDown,
		},
	}

	if app.Run(args) != nil {
		panic(err)
	}
	return nil
}

func serverAction(cliCtx *cli.Context) error {
	svr := server.NewServer(&cfg.Server)

	svc, err := service.NewService(cfg)
	if err != nil {
		return err
	}

	if err = svr.Register(svc); err != nil {
		logger.Error(err, "register server failed")
		return err
	}

	logger.Info("starting http server..", "http", cfg.Server.HTTP)
	if err = svr.Serve(); err != nil {
		logger.Error(err, "start server failed")
		return err
	}
	return nil
}

func migrateUp(cliCtx *cli.Context) error {
	m, err := migrateV4.New("file://migrations", cfg.MySQL.String())
	if err != nil {
		logger.Error(err, "error create migration", err.Error())
		return err
	}

	if err = m.Up(); err != nil && err != migrateV4.ErrNoChange {
		logger.Error(err, "error when migration up", err.Error())
		return err
	}
	logger.Info("migrate up success!")
	return nil
}

func migrateDown(cliCtx *cli.Context) error {
	m, err := migrateV4.New("file://migrations", cfg.MySQL.String())
	if err != nil {
		logger.Error(err, "error create migration", err.Error())
		return err
	}

	if err = m.Steps(-1); err != nil {
		logger.Error(err, "error when migration down", err.Error())
		return err
	}
	logger.Info("migrate down success!")
	return nil
}
