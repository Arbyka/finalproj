package controller

import (
    "net/http"
    "project-root/config"
    "project-root/entity"
    "github.com/gin-gonic/gin"
)

func SalesReport(c *gin.Context) {
    var products []entity.Product
    config.DB.Order("stock DESC").Find(&products)
    c.JSON(http.StatusOK, products)
}
