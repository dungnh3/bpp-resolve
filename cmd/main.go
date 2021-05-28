package main

import (
	"github.com/dungnh3/bpp-resolve/config"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/dungnh3/bpp-resolve/internal/server"
	"github.com/dungnh3/bpp-resolve/internal/service"
	"github.com/go-logr/logr"
	"github.com/urfave/cli"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
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
			Name:   "auto-migrate",
			Usage:  "auto migration",
			Action: migrate,
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

func migrate(cliCtx *cli.Context) error {
	db, err := gorm.Open(mysql.Open(cfg.MySQL.DSN()), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.Debug().AutoMigrate(
		&model.Wager{},
		&model.Purchase{},
	)
	if err != nil {
		return err
	}
	return nil
}
