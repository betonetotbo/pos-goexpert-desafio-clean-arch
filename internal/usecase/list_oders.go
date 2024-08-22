package usecase

import (
	"context"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database/repository"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/entity"
)

type ListOrdersInputDTO struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type OrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersOutputDTO struct {
	Orders []OrderDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	repo entity.OrderRepository
}

func NewListOrdersUseCase(ctx context.Context) *ListOrdersUseCase {
	db := database.GetDB(ctx)
	return &ListOrdersUseCase{
		repo: repository.NewOrderRepository(db),
	}
}

func (c *ListOrdersUseCase) Execute(input *ListOrdersInputDTO) (*ListOrdersOutputDTO, error) {
	orders, err := c.repo.List(input.Limit, input.Offset)
	if err != nil {
		return nil, err
	}
	orderDTOs := make([]OrderDTO, len(orders))
	for i, order := range orders {
		orderDTOs[i] = OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}
	return &ListOrdersOutputDTO{
		Orders: orderDTOs,
	}, nil
}
