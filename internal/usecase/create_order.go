package usecase

import (
	"context"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database/repository"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/entity"
	"github.com/google/uuid"
)

type CreateOrderInputDTO struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	repo entity.OrderRepository
}

func NewCreateOrderUseCase(ctx context.Context) *CreateOrderUseCase {
	db := database.GetDB(ctx)
	return &CreateOrderUseCase{
		repo: repository.NewOrderRepository(db),
	}
}

func (c *CreateOrderUseCase) Execute(input *CreateOrderInputDTO) (*CreateOrderOutputDTO, error) {
	order := entity.Order{
		ID:    uuid.New().String(),
		Price: input.Price,
		Tax:   input.Tax,
	}
	err := order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	if err := c.repo.Save(&order); err != nil {
		return nil, err
	}
	dto := &CreateOrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.Price + order.Tax,
	}
	return dto, nil
}
