package service

import (
	"github.com/dungnh3/bpp-resolve/config"
	"github.com/dungnh3/bpp-resolve/internal/domain/repository"
	"github.com/dungnh3/bpp-resolve/internal/usecase"
	"github.com/dungnh3/bpp-resolve/pkg/database"
	"github.com/go-logr/logr"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	logger logr.Logger
	uc     *usecase.UseCase
}

func NewService(cfg *config.Config) (*Service, error) {
	logger := cfg.Logger.MustBuildLogR()
	db := mustConnectMySQL(&cfg.MySQL, logger)
	repo := repository.NewRepository(db, logger)
	uc := usecase.NewUseCase(logger, repo)

	svc := &Service{
		logger: logger,
		uc:     uc,
	}
	return svc, nil
}

func mustConnectMySQL(cfg *database.MySQLConfig, logger logr.Logger) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxIdleTime(60 * time.Second)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	err = db.Raw("SELECT 1").Error
	if err != nil {
		panic(err)
	}
	logger.Info("connect database success", "host", cfg.Host, "port", cfg.Port, "database", cfg.Database)
	return db
}
