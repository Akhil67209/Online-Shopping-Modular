package main

import (
	"github.com/Akhil67209/Online-Shopping-Modular/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Routes
	router.GET("/api/products", controllers.GetProducts)                                   // list products with filters
	router.GET("/api/products/:id", controllers.GetProductDetail)                          // product detail by ID
	router.GET("/api/products/:id/recommendations", controllers.GetProductRecommendations) // recommendations

	router.Run("0.0.0.0:8080") // starts server at localhost:8080
}
