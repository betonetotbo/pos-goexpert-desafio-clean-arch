package graph

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUC *usecase.CreateOrderUseCase
	ListOrdersUC  *usecase.ListOrdersUseCase
}
