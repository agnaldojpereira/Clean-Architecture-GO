// internal/repository/order_repository.go
package repository

import (
	"context"
	"database/sql"

	"github.com/agnaldojpereira/Clean-Architecture-GO/internal/model"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) ListOrders(ctx context.Context) ([]model.Order, error) {
	query := `SELECT id, customer_id, total_amount, status, created_at FROM orders`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var o model.Order
		if err := rows.Scan(&o.ID, &o.CustomerID, &o.TotalAmount, &o.Status, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}
