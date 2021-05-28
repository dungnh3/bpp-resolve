package model

import "time"

type Purchase struct {
	ID          uint32    `json:"id" gorm:"primaryKey;autoIncrement"`
	WagerID     uint32    `json:"wager_id"`
	BuyingPrice float64   `json:"buying_price"`
	BoughtAt    time.Time `json:"bought_at" gorm:"<-:create;autoCreateTime"`
}
