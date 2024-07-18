package transform

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"time"
)

func OrderToEntity(order *pb.Order) (*database.Order, error) {
	items := make([]database.OrderItem, len(order.Items))
	for i, item := range order.Items {
		items[i] = OrderItemToEntity(item)
	}
	date, err := time.Parse(time.RFC3339, order.Date)
	if err != nil {
		return nil, err
	}
	return &database.Order{
		ID:       order.Id,
		Customer: order.Customer,
		Date:     date,
		Total:    order.Total,
		Items:    items,
	}, nil
}

func OrderItemToEntity(item *pb.OrderItem) database.OrderItem {
	return database.OrderItem{
		ID:       item.Id,
		Product:  item.Product,
		Price:    item.Price,
		Quantity: int(item.Quantity),
		Total:    item.Total,
	}
}
