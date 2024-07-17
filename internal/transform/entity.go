package transform

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
)

func OrderToEntity(order *pb.Order) *database.Order {
	items := make([]database.OrderItem, len(order.Items))
	for i, item := range order.Items {
		items[i] = OrderItemToEntity(item)
	}
	return &database.Order{
		ID:       order.Id,
		Customer: order.Customer,
		Date:     order.Date.AsTime(),
		Total:    order.Total,
		Items:    items,
	}
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
