package services

import (
	"irctc/models"
)

type ProductFilter struct {
	Category string
	Brand    string
	Size     string
	SortBy   string
	Page     int
	Limit    int
}

func FetchFilteredProducts(filter ProductFilter) ([]models.Product, error) {
	return models.QueryProductsFromDB(filter)
}

func FetchProductByID(id string) (models.Product, error) {
	return models.FindProductByID(id)
}

func GetRecommendationsForProduct(id string) []models.Product {
	return models.GetTopTrendingProducts()
}
