package product

type InsertProductRequest struct {
	CategoryId string  `json:"categoryId"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
}
