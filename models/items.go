package models

import "time"

type Item struct {
	ItemId      uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;unique" `
	Description string
	Quantity    uint
	OrderId     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
