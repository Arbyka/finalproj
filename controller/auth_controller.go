package controller

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"

    "project-root/config"
    "project-root/entity"
)

var jwtSecret = []byte("your_secret_key")

func Register(c *gin.Context) {
    var user entity.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if user.Email == "" || user.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
        return
    }

    user.Password = string(hashedPassword)
    if user.Role != "admin" && user.Role != "customer" {
        user.Role = "customer"
    }

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var loginData entity.User
    if err := c.ShouldBindJSON(&loginData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user entity.User
    if err := config.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // ===> Tambahkan role ke dalam token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "role":    user.Role,    // <= Tambahkan ini
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    tokenString, _ := token.SignedString(jwtSecret)

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}