package graph

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service pb.OrderServiceServer
	DB      *gorm.DB
}
