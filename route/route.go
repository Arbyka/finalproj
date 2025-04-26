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
        product.GET("", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), controller.GetProducts)
        product.POST("", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), controller.CreateProduct)
        product.PUT("/:id", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), controller.UpdateProduct)
        product.DELETE("/:id", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), controller.DeleteProduct)
    }

    r.POST("/orders", controller.CreateOrder)
    r.GET("/orders", controller.GetAllOrders)
    r.GET("/orders/:id", controller.GetOrderByID)
    r.PUT("/orders/:id/status", controller.UpdateOrderStatus)

    r.GET("/report", middleware.JWTAuthMiddleware(), controller.SalesReport)
    r.POST("/payment", middleware.JWTAuthMiddleware(), controller.DummyPayment)

    r.PUT("/orders/:id/confirm", controller.ConfirmOrder)

    r.POST("/product-images", middleware.JWTAuthMiddleware(), middleware.AdminOnlyMiddleware(), controller.CreateProductImage)

    return r
}
