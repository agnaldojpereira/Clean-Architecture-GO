package model

import "time"

// Order representa a estrutura de um pedido no sistema
type Order struct {
	ID          int64     `json:"id"`           // ID único do pedido
	CustomerID  int64     `json:"customer_id"`  // ID do cliente que fez o pedido
	TotalAmount float64   `json:"total_amount"` // Valor total do pedido
	Status      string    `json:"status"`       // Status atual do pedido (ex: "pendente", "entregue")
	CreatedAt   time.Time `json:"created_at"`   // Data e hora de criação do pedido
}
