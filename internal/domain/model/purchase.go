package model

import (
	"github.com/shopspring/decimal"
	"time"
)

type Purchase struct {
	ID          uint32          `json:"id" gorm:"primaryKey;autoIncrement"`
	WagerID     uint32          `json:"wager_id"`
	BuyingPrice decimal.Decimal `json:"buying_price" gorm:"type:decimal(10,2)"`
	BoughtAt    *time.Time      `json:"bought_at" gorm:"<-:create;autoCreateTime"`
}
