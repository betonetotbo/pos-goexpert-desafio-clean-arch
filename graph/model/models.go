package model

type Order struct {
	ID       string  `json:"id"`
	Customer string  `json:"customer"`
	Date     string  `json:"date"`
	Total    float64 `json:"total"`
}

type OrderItem struct {
	ID       string  `json:"id"`
	Product  string  `json:"product"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Total    float64 `json:"total"`
}
