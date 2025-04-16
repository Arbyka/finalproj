package route

import (
    "github.com/gin-gonic/gin"
    "project-root/controller"
    "project-root/middleware"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    r.POST("/register", controller.Register)
    r.POST("/login", controller.Login)

    product := r.Group("/products")
    {
        product.GET("", controller.GetProducts)
        product.POST("", middleware.JWTAuthMiddleware(), controller.CreateProduct)
        product.PUT("/:id", middleware.JWTAuthMiddleware(), controller.UpdateProduct)
        product.DELETE("/:id", middleware.JWTAuthMiddleware(), controller.DeleteProduct)
    }

    r.GET("/report", middleware.JWTAuthMiddleware(), controller.SalesReport)
    r.POST("/payment", middleware.JWTAuthMiddleware(), controller.DummyPayment)

    return r
}
