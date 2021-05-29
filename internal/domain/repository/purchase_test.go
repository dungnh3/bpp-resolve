package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/shopspring/decimal"
)

func (suite *TestRepositorySuite) TestRepository_RecordPurchase() {
	purchase := &model.Purchase{
		WagerID:     1,
		BuyingPrice: decimal.NewFromInt(100),
	}

	sql := "INSERT INTO `purchases` (`wager_id`,`buying_price`,`bought_at`) VALUES (?,?,?)"
	suite.mock.ExpectExec(sql).
		WithArgs(purchase.WagerID, purchase.BuyingPrice, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	err := suite.repo.RecordPurchase(context.Background(), purchase)
	if err != nil {
		suite.Failf("FAILED", "error should be nil, but got: ", err)
	}
}
