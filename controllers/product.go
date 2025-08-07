package controllers

import (
	"net/http"
	"strconv"

	"github.com/Akhil67209/Online-Shopping-Modular/services"

	"github.com/gin-gonic/gin"
)

// GetProducts handles listing of products with filters
func GetProducts(c *gin.Context) {
	filters := services.ProductFilter{
		Category: c.DefaultQuery("category", "mens-jeans"),
		Brand:    c.Query("brand"),
		Size:     c.Query("size"),
		SortBy:   c.DefaultQuery("sort_by", "popularity"),
		Page:     atoiOrDefault(c.DefaultQuery("page", "1"), 1),
		Limit:    atoiOrDefault(c.DefaultQuery("limit", "20"), 20),
	}

	products, err := services.FetchFilteredProducts(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductDetail returns product details by ID
func GetProductDetail(c *gin.Context) {
	id := c.Param("id")
	product, err := services.FetchProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// GetProductRecommendations returns dummy recommendations
func GetProductRecommendations(c *gin.Context) {
	id := c.Param("id")
	recommendations := services.GetRecommendationsForProduct(id)
	c.JSON(http.StatusOK, recommendations)
}

func atoiOrDefault(s string, def int) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return val
}
