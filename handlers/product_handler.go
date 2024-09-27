package handlers

import (
    "database/sql"
    "net/http"
    "github.com/gin-gonic/gin"
    "go-store-api/models"
)

// Получение всех продуктов
func GetAllProducts(c *gin.Context, db *sql.DB) {
    rows, err := db.Query("SELECT id, name, description, price, category, stock FROM products")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Stock); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            return
        }
        products = append(products, product)
    }

    c.JSON(http.StatusOK, products)
}

// Получение информации о продукте по ID
func GetProductByID(c *gin.Context, db *sql.DB) {
    id := c.Param("id")
    var product models.Product
    err := db.QueryRow("SELECT id, name, description, price, category, stock FROM products WHERE id = ?", id).
        Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category, &product.Stock)

    if err == sql.ErrNoRows {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    } else if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// Обновление продукта (только администратор)
func UpdateProduct(c *gin.Context, db *sql.DB) {
    id := c.Param("id")
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
        return
    }

    _, err := db.Exec("UPDATE products SET name=?, description=?, price=?, category=?, stock=? WHERE id=?", product.Name, product.Description, product.Price, product.Category, product.Stock, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
