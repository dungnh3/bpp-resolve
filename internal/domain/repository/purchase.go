package repository

import (
	"context"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
)

type PurchaseRepository interface {
	RecordPurchase(ctx context.Context, purchase *model.Purchase) error
}

func (r *Repository) RecordPurchase(ctx context.Context, purchase *model.Purchase) error {
	tx := r.db.WithContext(ctx).Create(purchase)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}
