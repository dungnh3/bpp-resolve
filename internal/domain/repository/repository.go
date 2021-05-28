package repository

import (
	"context"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"gorm.io/gorm"
)

type IRepository interface {
	RecordWager(ctx context.Context, wager *model.Wager) error
	FindWagerByID(ctx context.Context, id uint32) (*model.Wager, error)
	UpdateWagerByID(ctx context.Context, wager *model.Wager) error
	ListWager(ctx context.Context, offset, limit int) ([]*model.Wager, error)

	RecordPurchase(ctx context.Context, purchase *model.Purchase) error
}

type Repository struct {
	db *gorm.DB
}

func (r *Repository) FindWagerByID(ctx context.Context, id uint32) (*model.Wager, error) {
	panic("implement me")
}

func (r *Repository) UpdateWagerByID(ctx context.Context, wager *model.Wager) error {
	tx := r.db.WithContext(ctx).Model(wager).Updates(model.Wager{
		CurrentSellingPrice: wager.CurrentSellingPrice,
		PercentageSold:      wager.PercentageSold,
		AmountSold:          wager.AmountSold,
	})

	if err := tx.Error; err != nil {
		return err
	}

	if rowsAffected := tx.RowsAffected; rowsAffected != 1 {
		return ErrRecordAffectedNotExpected
	}
	return nil
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) RecordWager(ctx context.Context, wager *model.Wager) error {
	tx := r.db.WithContext(ctx).Create(wager)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) ListWager(ctx context.Context, offset, limit int) ([]*model.Wager, error) {
	var wagers []*model.Wager
	tx := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&wagers)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return wagers, nil
}

func (r *Repository) RecordPurchase(ctx context.Context, purchase *model.Purchase) error {
	tx := r.db.WithContext(ctx).Create(purchase)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
