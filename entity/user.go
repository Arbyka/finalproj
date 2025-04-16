package entity

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string
	Email     string    `gorm:"unique"`
	Password  string
	Role      string    `gorm:"type:enum('admin','customer')"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
