package handler

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/seu-usuario/projeto-pedidos/internal/model"
	"github.com/seu-usuario/projeto-pedidos/internal/service"
)

// Esta struct serve como o resolver raiz
type Resolver struct {
	orderService *service.OrderService
}

// Query fornece os resolvers para as queries GraphQL
type Query struct{ *Resolver }

// NewGraphQLHandler cria uma nova instância do handler GraphQL
func NewGraphQLHandler(orderService *service.OrderService) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			orderService: orderService,
		},
	})
}

// ListOrders resolve a query para listar todos os pedidos
func (r *Query) ListOrders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.orderService.ListOrders()
	if err != nil {
		return nil, err
	}

	var orderPtrs []*model.Order
	for i := range orders {
		orderPtrs = append(orderPtrs, &orders[i])
	}

	return orderPtrs, nil
}
