package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"project-root/entity"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:Semangat45.@tcp(127.0.0.1:3306)/goproject?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi ke database:", err)
	}

	DB = database
	fmt.Println("Database connected")

	// Migrasi model
	err = DB.AutoMigrate(
		&entity.User{},
		&entity.Product{},
		&entity.Order{},
		&entity.OrderItem{},
		&entity.Transaction{},
		&entity.SalesReport{}, // opsional
	)

	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}
}
