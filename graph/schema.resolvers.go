package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/usecase"
	"math"

	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/graph/model"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input model.CreateOrderInput) (*model.Order, error) {
	order, err := r.CreateOrderUC.Execute(&usecase.CreateOrderInputDTO{
		Price: input.Price,
		Tax:   input.Tax,
	})
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}

// ListOrders is the resolver for the listOrders field.
func (r *queryResolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	list, err := r.ListOrdersUC.Execute(&usecase.ListOrdersInputDTO{
		Limit:  math.MaxInt32,
		Offset: 0,
	})
	if err != nil {
		return nil, err
	}

	result := make([]*model.Order, len(list.Orders))
	for i, order := range list.Orders {
		result[i] = &model.Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}
	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
