package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/dungnh3/bpp-resolve/internal/domain/model"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
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

func (suite *TestRepositorySuite) TestRepository_RecordWagerPriceByID() {
	buyingPrice := decimal.NewFromInt(100)
	sellingPrice := decimal.NewFromInt(200)

	sql := "UPDATE `wagers` SET `amount_sold`=IFNULL(amount_sold, 0) + ?,`current_selling_price`=current_selling_price - ?,`percentage_sold`=IFNULL(percentage_sold, 0) + ? WHERE id = ?"
	suite.mock.ExpectExec(sql).
		WithArgs(buyingPrice, buyingPrice, sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repo.RecordWagerPriceByID(context.Background(), 1, buyingPrice, sellingPrice)
	if err != nil {
		suite.Failf("FAILED", "error should be nil, but got: ", err)
	}
}

func (suite *TestRepositorySuite) TestRepository_FindWagerByID() {
	wager := &model.Wager{
		ID:                  1,
		TotalWagerValue:     1000,
		Odds:                30,
		SellingPercentage:   50,
		SellingPrice:        decimal.NewFromInt(800),
		CurrentSellingPrice: decimal.NewFromInt(600),
		PercentageSold: func() *decimal.Decimal {
			result := decimal.NewFromInt(25)
			return &result
		}(),
		AmountSold: func() *decimal.Decimal {
			result := decimal.NewFromInt(200)
			return &result
		}(),
		PlacedAt: nil,
	}

	sql := "SELECT * FROM `wagers` WHERE id = ?"
	suite.mock.ExpectQuery(sql).
		WithArgs(wager.ID).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "total_wager_value", "odds", "selling_percentage", "selling_price",
					"current_selling_price", "percentage_sold", "amount_sold", "placed_at"},
			).AddRow(wager.ID, wager.TotalWagerValue, wager.Odds, wager.SellingPercentage, wager.SellingPrice,
				wager.CurrentSellingPrice, wager.PercentageSold, wager.AmountSold, wager.PlacedAt),
		)

	resp, err := suite.repo.FindWagerByID(context.Background(), wager.ID)
	if err != nil {
		suite.Failf("FAILED", "error should be nil, but got: ", err)
	}
	assert.Equal(suite.T(), wager.SellingPrice, resp.SellingPrice)
	assert.Equal(suite.T(), wager.TotalWagerValue, resp.TotalWagerValue)
	assert.Equal(suite.T(), wager.Odds, resp.Odds)
	assert.Equal(suite.T(), wager.AmountSold, resp.AmountSold)
	assert.Equal(suite.T(), wager.PercentageSold, resp.PercentageSold)
	assert.Equal(suite.T(), wager.SellingPercentage, resp.SellingPercentage)
}

func (suite *TestRepositorySuite) TestRepository_RecordPurchasingWagerByID() {
	currentSellingPrice := decimal.NewFromInt(900)
	amountSold := decimal.NewFromInt(300)
	percentageSold := decimal.NewFromInt(25)

	sql := "UPDATE `wagers` SET `amount_sold`=?,`current_selling_price`=?,`percentage_sold`=? WHERE id = ?"
	suite.mock.ExpectExec(sql).
		WithArgs(amountSold, currentSellingPrice, percentageSold, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := suite.repo.RecordPurchasingWagerByID(context.Background(), 1, currentSellingPrice, amountSold, percentageSold)
	if err != nil {
		suite.Failf("FAILED", "error should be nil, but got: ", err)
	}
}

func (suite *TestRepositorySuite)TestRepository_SelectForUpdateWagerByID() {
	wager := &model.Wager{
		ID:                  1,
		TotalWagerValue:     1000,
		Odds:                30,
		SellingPercentage:   50,
		SellingPrice:        decimal.NewFromInt(800),
		CurrentSellingPrice: decimal.NewFromInt(600),
		PercentageSold: func() *decimal.Decimal {
			result := decimal.NewFromInt(25)
			return &result
		}(),
		AmountSold: func() *decimal.Decimal {
			result := decimal.NewFromInt(200)
			return &result
		}(),
		PlacedAt: nil,
	}

	sql := "SELECT * FROM `wagers` WHERE `wagers`.`id` = ? ORDER BY `wagers`.`id` LIMIT 1 FOR UPDATE"
	suite.mock.ExpectQuery(sql).
		WithArgs(wager.ID).
		WillReturnRows(
			sqlmock.NewRows(
				[]string{"id", "total_wager_value", "odds", "selling_percentage", "selling_price",
					"current_selling_price", "percentage_sold", "amount_sold", "placed_at"},
			).AddRow(wager.ID, wager.TotalWagerValue, wager.Odds, wager.SellingPercentage, wager.SellingPrice,
				wager.CurrentSellingPrice, wager.PercentageSold, wager.AmountSold, wager.PlacedAt),
		)

	resp, err := suite.repo.SelectForUpdateWagerByID(context.Background(), wager.ID)
	if err != nil {
		suite.Failf("FAILED", "error should be nil, but got: ", err)
	}
	assert.Equal(suite.T(), wager.SellingPrice, resp.SellingPrice)
	assert.Equal(suite.T(), wager.TotalWagerValue, resp.TotalWagerValue)
	assert.Equal(suite.T(), wager.Odds, resp.Odds)
	assert.Equal(suite.T(), wager.AmountSold, resp.AmountSold)
	assert.Equal(suite.T(), wager.PercentageSold, resp.PercentageSold)
	assert.Equal(suite.T(), wager.SellingPercentage, resp.SellingPercentage)
}