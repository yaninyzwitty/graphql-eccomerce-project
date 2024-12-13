package model

import "time"

type Customer struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	CreatedAt string   `json:"created_at"`
	Orders    []*Order `json:"orders"`
}

type DbCustomer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	Orders    []*Order  `json:"orders"`
}
