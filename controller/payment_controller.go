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

    // Bind JSON dari body
    if err := c.ShouldBindJSON(&paymentReq); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Cari order berdasarkan ID
    var order entity.Order
    if err := config.DB.First(&order, paymentReq.OrderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    // Update status jadi "paid"
    order.Status = "paid"
    if err := config.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order status"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":         "Payment processed successfully",
        "payment_method": paymentReq.PaymentMethod,
        "order_id":       paymentReq.OrderID,
    })
}
