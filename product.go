package models

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Brand    string  `json:"brand"`
	Price    float64 `json:"price"`
	Size     string  `json:"size"`
	Rating   float64 `json:"rating"`
	Stock    int     `json:"stock"`
	ImageURL string  `json:"image_url"`
}

type ProductFilter struct {
	Category string
	Brand    string
	Size     string
	SortBy   string
	Page     int
	Limit    int
}

func QueryProductsFromDB(filter interface{}) ([]Product, error) {
	// Dummy data
	return []Product{
		{ID: "1", Name: "Levis Slim Fit Jeans", Brand: "Levis", Price: 1999, Size: "32", Rating: 4.2, Stock: 20, ImageURL: "https://example.com/img1.jpg"},
		{ID: "2", Name: "Roadster Regular Fit", Brand: "Roadster", Price: 1499, Size: "34", Rating: 4.0, Stock: 12, ImageURL: "https://example.com/img2.jpg"},
	}, nil
}

func FindProductByID(id string) (Product, error) {
	return Product{ID: id, Name: "Sample Product", Price: 1299, Size: "32", Brand: "Levis", Rating: 4.5, Stock: 10, ImageURL: "https://example.com/sample.jpg"}, nil
}

func GetTopTrendingProducts() []Product {
	return []Product{
		{ID: "10", Name: "Trending Jeans 1", Price: 1699},
		{ID: "11", Name: "Trending Jeans 2", Price: 1899},
	}
}
