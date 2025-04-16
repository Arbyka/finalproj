package entity

type SalesReport struct {
	ID           uint    `gorm:"primaryKey"`
	ProductID    uint
	Product      Product
	TotalSold    int
	TotalRevenue float64
	Period       string
}
