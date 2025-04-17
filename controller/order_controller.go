package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"project-root/entity"
	"project-root/config"
)

// Struct untuk input dari user
type CreateOrderInput struct {
	UserID          uint `json:"user_id"`
	ShippingAddress string `json:"shipping_address"`
	Items           []struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	} `json:"items"`
}

// Create Order
func CreateOrder(c *gin.Context) {
	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var orderItems []entity.OrderItem
	var totalPrice float64

	for _, item := range input.Items {
		var product entity.Product
		if err := config.DB.First(&product, item.ProductID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		if product.Stock < item.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock for product ID " + strconv.Itoa(int(item.ProductID))})
			return
		}

		// Kurangi stok produk
		product.Stock -= item.Quantity
		config.DB.Save(&product)

		orderItems = append(orderItems, entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})

		totalPrice += product.Price * float64(item.Quantity)
	}

	order := entity.Order{
		UserID:          input.UserID,
		ShippingAddress: input.ShippingAddress,
		Status:          "pending",
		TotalPrice:      totalPrice,
		OrderItems:      orderItems,
	}

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

// Get All Orders
func GetAllOrders(c *gin.Context) {
	var orders []entity.Order
	err := config.DB.Preload("OrderItems").Preload("OrderItems.Product").Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Get Order by ID
func GetOrderByID(c *gin.Context) {
	id := c.Param("id")
	var order entity.Order

	err := config.DB.Preload("OrderItems").Preload("OrderItems.Product").First(&order, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve order"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Update Order Status
func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var order entity.Order

	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.Status = input.Status
	config.DB.Save(&order)

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated", "order": order})
}

func ConfirmOrder(c *gin.Context) {
    orderIDStr := c.Param("id")
    orderID, err := strconv.Atoi(orderIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order entity.Order
    if err := config.DB.First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    if order.Status != "paid" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Order must be paid before shipping"})
        return
    }

    order.Status = "shipped"
    config.DB.Save(&order)

    c.JSON(http.StatusOK, gin.H{
        "message": "Order confirmed and marked for shipment",
        "order":   order,
    })
}