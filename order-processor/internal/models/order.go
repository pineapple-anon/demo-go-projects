package models

type Order struct {
	ID        int64   `json:"id"`
	Symbol  string  `json:"symbol"`
	Price     float64 `json:"price"`
	Quantity  float64 `json:"quantity"`
	Timestamp int64   `json:"timestamp"`
}