package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func AdminOnlyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        roleInterface, exists := c.Get("role")
        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found"})
            c.Abort()
            return
        }

        role, ok := roleInterface.(string)
        if !ok || role != "admin" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admins only"})
            c.Abort()
            return
        }

        c.Next()
    }
}
