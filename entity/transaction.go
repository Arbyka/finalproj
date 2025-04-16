package entity

import "time"

type Transaction struct {
	ID         uint      `gorm:"primaryKey"`
	OrderID    uint
	Order      Order
	Amount     float64
	Status     string    `gorm:"type:enum('success','failed')"`
	Method     string
	PaymentDate time.Time
}
