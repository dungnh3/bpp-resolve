package repository

import (
	"context"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WagerRepository interface {
	InitializeWager(ctx context.Context, wager *model.Wager) error
	FindWagerByID(ctx context.Context, id uint32) (*model.Wager, error)
	RecordWagerPriceByID(ctx context.Context, wagerId uint32, buyingPrice, sellingPrice float64) error
	FindWagers(ctx context.Context, offset, limit int) ([]*model.Wager, error)
}

func (r *Repository) InitializeWager(ctx context.Context, wager *model.Wager) error {
	tx := r.db.WithContext(ctx).Create(wager)
	if err := tx.Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) FindWagers(ctx context.Context, offset, limit int) ([]*model.Wager, error) {
	var wagers []*model.Wager
	tx := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&wagers)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return wagers, nil
}

func (r *Repository) FindWagerByID(ctx context.Context, id uint32) (*model.Wager, error) {
	panic("implement me")
}

func (r *Repository) RecordWagerPriceByID(ctx context.Context, wagerId uint32, buyingPrice float64, sellingPrice float64) error {
	tx := r.db.WithContext(ctx).Model(&model.Wager{}).
		Where("id = ?", wagerId).
		Updates(map[string]interface{}{
			"current_selling_price": gorm.Expr("? - ?",
				clause.Column{Table: clause.CurrentTable, Name: "current_selling_price"},
				buyingPrice,
			),
			"amount_sold": gorm.Expr("? + ?",
				clause.Column{Table: clause.CurrentTable, Name: "amount_sold"},
				buyingPrice,
			),
			"percentage_sold": gorm.Expr("? + ?",
				clause.Column{Table: clause.CurrentTable, Name: "percentage_sold"},
				100*buyingPrice/sellingPrice,
			),
		})

	if err := tx.Error; err != nil {
		return err
	}

	if rowsAffected := tx.RowsAffected; rowsAffected != 1 {
		return ErrRecordAffectedNotExpected
	}
	return nil
}
