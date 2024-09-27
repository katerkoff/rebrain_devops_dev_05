package handlers

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
)

// Добавление товара в корзину
func AddToCart(c *gin.Context, db *sql.DB) {
    var cartItem struct {
        ProductID int `json:"productId"`
        Quantity  int `json:"quantity"`
    }

    if err := c.ShouldBindJSON(&cartItem); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    if cartItem.Quantity == 0 {
        cartItem.Quantity = 1
    }

    _, err := db.Exec("INSERT INTO cart (product_id, quantity) VALUES (?, ?)", cartItem.ProductID, cartItem.Quantity)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}
