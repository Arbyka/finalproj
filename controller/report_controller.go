package controller

import (
    "net/http"
    "project-root/config"
    // "project-root/entity"
    "github.com/gin-gonic/gin"
)

func SalesReport(c *gin.Context) {
    statusFilter := c.Query("status")

    type ReportResult struct {
        ProductID    uint    `json:"product_id"`
        ProductName  string  `json:"product_name"`
        TotalSold    int     `json:"total_sold"`
        TotalRevenue float64 `json:"total_revenue"`
    }

    var results []ReportResult
    query := config.DB.Table("order_items").
        Select("products.id as product_id, products.name as product_name, SUM(order_items.quantity) as total_sold, SUM(order_items.quantity * products.price) as total_revenue").
        Joins("JOIN products ON order_items.product_id = products.id").
        Joins("JOIN orders ON order_items.order_id = orders.id").
        Group("products.id, products.name")

    if statusFilter != "" {
        query = query.Where("orders.status = ?", statusFilter)
    }

    if err := query.Scan(&results).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, results)
}
