package middleware

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Мидлварь для проверки Bearer токена администратора
func AdminAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != "Bearer admin-token" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}

// Мидлварь для проверки Bearer токена пользователя
func UserAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != "Bearer user-token" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}
