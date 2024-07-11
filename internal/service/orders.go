package service

import "github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
