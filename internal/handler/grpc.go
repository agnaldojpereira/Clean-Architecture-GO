package handler

import (
	"context"

	"github.com/seu-usuario/projeto-pedidos/internal/pb"
	"github.com/seu-usuario/projeto-pedidos/internal/service"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GRPCHandler lida com as requisições gRPC
type GRPCHandler struct {
	pb.UnimplementedOrderServiceServer
	orderService *service.OrderService
}

// NewGRPCHandler cria uma nova instância de GRPCHandler
func NewGRPCHandler(orderService *service.OrderService) *GRPCHandler {
	return &GRPCHandler{orderService: orderService}
}

// ListOrders implementa o método ListOrders do serviço gRPC
func (h *GRPCHandler) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := h.orderService.ListOrders()
	if err != nil {
		return nil, err
	}

	var pbOrders []*pb.Order
	for _, order := range orders {
		pbOrders = append(pbOrders, &pb.Order{
			Id:           order.ID,
			CustomerName: order.CustomerName,
			TotalAmount:  order.TotalAmount,
			Status:       order.Status,
			CreatedAt:    timestamppb.New(order.CreatedAt),
		})
	}

	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
