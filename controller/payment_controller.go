package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func DummyPayment(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "Payment processed successfully"})
}