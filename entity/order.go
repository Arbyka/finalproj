package entity

import "time"

type Order struct {
	ID              uint         `gorm:"primaryKey"`
	UserID          uint
	User            User
	Status          string       `gorm:"type:enum('pending','paid','cancelled','delivered')"`
	TotalPrice      float64
	ShippingAddress string
	OrderItems      []OrderItem
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
