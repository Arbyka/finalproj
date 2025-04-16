package controller

import (
    "project-root/config"
    "project-root/entity"
    "net/http"
    // "strconv"

    "github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
    var products []entity.Product
    name := c.Query("name")
    query := config.DB
    if name != "" {
        query = query.Where("name LIKE ?", "%"+name+"%")
    }
    query.Find(&products)
    c.JSON(http.StatusOK, products)
}

func CreateProduct(c *gin.Context) {
    var product entity.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Create(&product)
    c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product entity.Product
    if err := config.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    config.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    config.DB.Delete(&entity.Product{}, id)
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
