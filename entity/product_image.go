package entity

type ProductImage struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ProductID uint   `json:"product_id"`
	URL       string `json:"url"`

	Product Product `gorm:"foreignKey:ProductID" json:"-"`
}
