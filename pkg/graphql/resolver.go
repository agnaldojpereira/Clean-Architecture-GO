package graphql

import (
	"context"

	"github.com/agnaldojpereira/Clean-Architecture-GO/internal/model"
	"github.com/agnaldojpereira/Clean-Architecture-GO/internal/service"
)

type Resolver struct {
	orderService *service.OrderService
}

func NewResolver(orderService *service.OrderService) *Resolver {
	return &Resolver{orderService: orderService}
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type QueryResolver interface {
	ListOrders(ctx context.Context) ([]*model.Order, error)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.orderService.ListOrders(ctx)
	if err != nil {
		return nil, err
	}
	
	// Convert []model.Order to []*model.Order
	orderPtrs := make([]*model.Order, len(orders))
	for i := range orders {
		orderPtrs[i] = &orders[i]
	}
	
	return orderPtrs, nil
}
