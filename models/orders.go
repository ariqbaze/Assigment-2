package models

import "time"

type Order struct {
	OrderId      uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Items        []Item
}
