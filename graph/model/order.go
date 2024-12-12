package model

type Order struct {
	ID        string       `json:"id"`
	Customer  *Customer    `json:"customer"`
	CreatedAt string       `json:"created_at"`
	Items     []*OrderItem `json:"items"`
}
