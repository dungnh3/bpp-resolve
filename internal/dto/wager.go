package dto

type CreateWagerDto struct {
	TotalWagerValue   uint32  `json:"total_wager_value"`
	Odds              uint32  `json:"odds"`
	SellingPercentage int     `json:"selling_percentage"`
	SellingPrice      float64 `json:"selling_price"`
}

type BuyingWagerDto struct {
	BuyingPrice float64 `json:"buying_price"`
}
