package model

type Customer struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	CreatedAt string   `json:"created_at"`
	Orders    []*Order `json:"orders"`
}
