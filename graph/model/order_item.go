package model

type OrderItem struct {
	ID        string   `json:"id"`
	Order     *Order   `json:"order"`
	Product   *Product `json:"product"`
	Quantity  int32    `json:"quantity"`
	Price     float64  `json:"price"`
	CreatedAt string   `json:"created_at"`
}
