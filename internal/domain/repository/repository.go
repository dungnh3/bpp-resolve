package repository

import (
	"github.com/go-logr/logr"
	"gorm.io/gorm"
)

type IRepository interface {
	WagerRepository
	PurchaseRepository

	Transaction(txFunc func(IRepository) error) error
}

type Repository struct {
	db     *gorm.DB
	logger logr.Logger
}

func NewRepository(db *gorm.DB, logger logr.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (r *Repository) WithTx(tx *gorm.DB) *Repository {
	newRepo := *r
	newRepo.db = tx
	return &newRepo
}

func (r *Repository) Transaction(txFunc func(IRepository) error) error {
	tx := r.db.Begin()
	defer func() {
		if rc := recover(); rc != nil {
			r.logger.Error(nil, "rollback now because listening recover: %v \n", rc)
			if execErr := tx.Rollback().Error; execErr != nil {
				r.logger.Error(execErr, "exception error when execute rollback")
			}
			panic(rc)
		}
	}()

	err := txFunc(r.WithTx(tx))
	if err != nil {
		if execErr := tx.Rollback().Error; execErr != nil {
			r.logger.Error(execErr, "exception error when execute rollback")
		}
		return err
	}
	return tx.Commit().Error
}
