package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/shopspring/decimal"
)

func (suite *TestRepositorySuite) TestRepository_InitializeWager() {
	wager := &model.Wager{
		TotalWagerValue:     1000,
		Odds:                40,
		SellingPercentage:   50,
		SellingPrice:        decimal.NewFromInt(600),
		CurrentSellingPrice: decimal.NewFromInt(600),
		PercentageSold:      nil,
		AmountSold:          nil,
	}

	sql := "INSERT INTO `wagers` (`total_wager_value`,`odds`,`selling_percentage`,`selling_price`,`current_selling_price`,`percentage_sold`,`amount_sold`,`placed_at`) VALUES (?,?,?,?,?,?,?,?)"
	suite.mock.ExpectExec(sql).
		WithArgs(wager.TotalWagerValue, wager.Odds, wager.SellingPercentage, wager.SellingPrice, wager.CurrentSellingPrice, wager.PercentageSold, wager.AmountSold, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewErrorResult(nil))

	err := suite.repo.InitializeWager(context.Background(), wager)
	if err != nil {
		suite.Failf("FAILED", "error should be nil, but got: ", err)
	}
}

