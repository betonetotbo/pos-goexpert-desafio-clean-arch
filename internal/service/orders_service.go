package service

import (
	"context"
	"fmt"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/transform"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"time"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{
		db: db,
	}
}

func (s *OrderService) ListOrders(_ context.Context, in *emptypb.Empty) (*pb.OrderList, error) {
	var orders []database.Order
	err := s.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	dtos := make([]*pb.Order, len(orders))
	for i, order := range orders {
		dtos[i] = transform.OrderToProtobuf(&order)
	}

	return &pb.OrderList{
		Orders: dtos,
	}, nil
}

func (s *OrderService) CreateOrder(_ context.Context, in *pb.Order) (*pb.Order, error) {
	if len(in.Items) == 0 {
		return nil, fmt.Errorf("É necessário informar um ou mais 'items'")
	}

	var total = 0.0
	for i, item := range in.Items {
		if item.Price <= 0.0 {
			return nil, fmt.Errorf("O item (%d / %s) não possui 'price' válido: %v", i+1, item.Product, item.Price)
		}
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("O item (%d / %s)não possui 'quantity' válida: %v", i+1, item.Product, item.Quantity)
		}

		item.Total = item.Price * float64(item.Quantity)
		total += item.Total
	}
	in.Total = total
	in.Date = timestamppb.New(time.Now())

	entity := transform.OrderToEntity(in)

	err := s.db.Create(entity).Error
	if err != nil {
		return nil, err
	}

	in.Id = entity.ID
	for i, item := range in.Items {
		item.Id = entity.Items[i].ID
		item.Total = entity.Total
	}

	return in, nil
}
