package model

import (
	"time"
)

type Wager struct {
	ID                  uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	TotalWagerValue     uint32    `json:"total_wager_value"`
	Odds                uint32    `json:"odds"`
	SellingPercentage   int       `json:"selling_percentage"`
	SellingPrice        float64   `json:"selling_price"`
	CurrentSellingPrice float64   `json:"current_selling_price"`
	PercentageSold      *int      `json:"percentage_sold"`
	AmountSold          *int      `json:"amount_sold"`
	PlacedAt            time.Time `json:"placed_at" gorm:"<-:create;autoCreateTime"`
}
