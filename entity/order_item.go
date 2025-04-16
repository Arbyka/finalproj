package entity

type OrderItem struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Product   Product
	Quantity  int
	Price     float64
}
