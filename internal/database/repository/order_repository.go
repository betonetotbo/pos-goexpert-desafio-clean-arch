package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/entity"
	"github.com/google/uuid"
)

type (
	orderRepository struct {
		db *sql.DB
	}
)

func NewOrderRepository(db *sql.DB) entity.OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r *orderRepository) Save(order *entity.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	q := database.New(tx)
	in := database.CreateOrderParams{
		ID:         uuid.New().String(),
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	err = q.CreateOrder(context.Background(), in)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("failed to rollback transaction: %v, root cause: %v", rbErr, err)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (r *orderRepository) List(limit, offset int32) ([]entity.Order, error) {
	q := database.New(r.db)

	orders, err := q.ListOrders(context.Background(), database.ListOrdersParams{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		return nil, err
	}

	result := make([]entity.Order, len(orders))
	for i, order := range orders {
		result[i] = entity.Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return result, nil
}
