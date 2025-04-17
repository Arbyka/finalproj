package controller

import (
    "net/http"
    "project-root/config"
    "project-root/entity"
    "github.com/gin-gonic/gin"
)

type PaymentRequest struct {
    OrderID       uint   `json:"order_id"`
    PaymentMethod string `json:"payment_method"`
}

func DummyPayment(c *gin.Context) {
    var paymentReq PaymentRequest

    if err := c.ShouldBindJSON(&paymentReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var order entity.Order
    if err := config.DB.Preload("OrderItems").First(&order, paymentReq.OrderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    // Simulasi kegagalan pembayaran
    if paymentReq.PaymentMethod == "fail" {
        // Update status pesanan jadi dibatalkan
        order.Status = "cancelled"
        config.DB.Save(&order)

        // Kembalikan stok produk
        for _, item := range order.OrderItems {
            var product entity.Product
            if err := config.DB.First(&product, item.ProductID).Error; err == nil {
                product.Stock += item.Quantity
                config.DB.Save(&product)
            }
        }

        c.JSON(http.StatusPaymentRequired, gin.H{
            "status": "Payment failed. Order cancelled and stock restored.",
        })
        return
    }

    // Jika sukses, ubah status jadi paid
    order.Status = "paid"
    config.DB.Save(&order)

    c.JSON(http.StatusOK, gin.H{
        "status":         "Payment processed successfully",
        "payment_method": paymentReq.PaymentMethod,
        "order_id":       paymentReq.OrderID,
    })
}
