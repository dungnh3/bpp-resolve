package dto

import "github.com/shopspring/decimal"

type CreateWagerDto struct {
	TotalWagerValue   uint32          `json:"total_wager_value"`
	Odds              uint32          `json:"odds"`
	SellingPercentage int             `json:"selling_percentage"`
	SellingPrice      decimal.Decimal `json:"selling_price"`
}

type BuyingWagerDto struct {
	BuyingPrice decimal.Decimal `json:"buying_price"`
}
