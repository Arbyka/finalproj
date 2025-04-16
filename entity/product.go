package entity

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string
	Category    string
	Price       float64
	Stock       int
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
