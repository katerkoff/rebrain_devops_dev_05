package main

import (
    "database/sql"
    "log"
    "go-store-api/handlers"
    "go-store-api/middleware"
    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Инициализация базы данных
    db, err := sql.Open("sqlite3", "./db/store.db")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Инициализация роутера
    r := gin.Default()

    // Маршруты для продуктов
    r.GET("/products", func(c *gin.Context) {
        handlers.GetAllProducts(c, db)
    })
    r.GET("/products/:id", func(c *gin.Context) {
        handlers.GetProductByID(c, db)
    })
    r.PUT("/products/:id", middleware.AdminAuth(), func(c *gin.Context) {
        handlers.UpdateProduct(c, db)
    })

    // Маршруты для корзины
    r.POST("/cart", middleware.UserAuth(), func(c *gin.Context) {
        handlers.AddToCart(c, db)
    })

    // Запуск сервера
    r.Run(":8080")
}
