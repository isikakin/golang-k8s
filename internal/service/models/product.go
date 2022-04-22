package models

type ProductDetail struct {
	Name     string       `json:"name"`
	Price    float64      `json:"price"`
	Category CategoryItem `json:"category"`
}

type CategoryItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
