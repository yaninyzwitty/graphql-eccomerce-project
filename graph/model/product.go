package model

import "time"

type Product struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Price      float64      `json:"price"`
	CreatedAt  string       `json:"created_at"`
	OrderItems []*OrderItem `json:"order_items"`
}
type DBProduct struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Price      float64      `json:"price"`
	CreatedAt  time.Time    `json:"created_at"`
	OrderItems []*OrderItem `json:"order_items"`
}
