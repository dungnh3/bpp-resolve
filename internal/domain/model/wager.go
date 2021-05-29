package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Wager struct {
	ID                  uint32           `json:"id" gorm:"primaryKey;autoIncrement"`
	TotalWagerValue     uint32           `json:"total_wager_value"`
	Odds                uint32           `json:"odds"`
	SellingPercentage   int              `json:"selling_percentage"`
	SellingPrice        decimal.Decimal  `json:"selling_price" gorm:"type:decimal(10,2)"`
	CurrentSellingPrice decimal.Decimal  `json:"current_selling_price" gorm:"type:decimal(10,2)"`
	PercentageSold      *decimal.Decimal `json:"percentage_sold" gorm:"type:decimal(10,2)"`
	AmountSold          *decimal.Decimal `json:"amount_sold" gorm:"type:decimal(10,2)"`
	PlacedAt            *time.Time       `json:"placed_at" gorm:"<-:create;autoCreateTime"`
}
