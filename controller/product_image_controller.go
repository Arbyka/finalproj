package controller

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"project-root/config"
	"project-root/entity"

	"github.com/gin-gonic/gin"
)

func CreateProductImage(c *gin.Context) {
	// Ambil form data
	productIDStr := c.PostForm("product_id")
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Pastikan folder uploads tersedia
	uploadDir := "./uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// Simpan file dengan nama unik
	fileName := fmt.Sprintf("img_%d_%d%s", productID, time.Now().Unix(), filepath.Ext(file.Filename))
	filePath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	// Simpan path ke database
	image := entity.ProductImage{
		ProductID: uint(productID),
		URL:       filePath,
	}

	if err := config.DB.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded", "data": image})
}
