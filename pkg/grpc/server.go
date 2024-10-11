// pkg/grpc/server.go
package grpc

import (
    "context"
    "github.com/agnaldojpereira/Clean-Architecture-GO/internal/service"
    "github.com/agnaldojpereira/Clean-Architecture-GO/pkg/grpc/proto"
    "google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
    pb.UnimplementedOrderServiceServer
    orderService *service.OrderService
}

func NewServer(orderService *service.OrderService) *Server {
    return &Server{orderService: orderService}
}

func (s *Server) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
    orders, err := s.orderService.ListOrders(ctx)
    if err != nil {
        return nil, err
    }

    var pbOrders []*pb.Order
    for _, o := range orders {
        pbOrders = append(pbOrders, &pb.Order{
            Id:          o.ID,
            CustomerId:  o.CustomerID,
            TotalAmount: o.TotalAmount,
            Status:      o.Status,
            CreatedAt:   timestamppb.New(o.CreatedAt),
        })
    }

    return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}