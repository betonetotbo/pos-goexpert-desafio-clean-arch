package service

import (
	"context"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	createOrders *usecase.CreateOrderUseCase
	listOrders   *usecase.ListOrdersUseCase
}

func NewOrderService(ctx context.Context) *OrderService {
	return &OrderService{
		createOrders: usecase.NewCreateOrderUseCase(ctx),
		listOrders:   usecase.NewListOrdersUseCase(ctx),
	}
}

func (s *OrderService) ListOrders(_ context.Context, in *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	list, err := s.listOrders.Execute(&usecase.ListOrdersInputDTO{
		Limit:  in.Limit,
		Offset: in.Offset,
	})
	if err != nil {
		return nil, err
	}

	dtos := make([]*pb.Order, len(list.Orders))
	for i, order := range list.Orders {
		dtos[i] = &pb.Order{
			Id:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return &pb.ListOrdersResponse{
		Orders: dtos,
	}, nil
}

func (s *OrderService) CreateOrder(_ context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order, err := s.createOrders.Execute(&usecase.CreateOrderInputDTO{
		Price: in.Price,
		Tax:   in.Tax,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
