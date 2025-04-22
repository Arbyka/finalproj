package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"project-root/entity"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Ambil dari environment
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Gunakan variabel untuk membentuk DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

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
		&entity.SalesReport{},
		&entity.ProductImage{},

	)

	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
	}
}
