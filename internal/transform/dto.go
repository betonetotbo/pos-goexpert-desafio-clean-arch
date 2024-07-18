package transform

import (
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/database"
	"github.com/betonetotbo/pos-goexpert-desafio-clean-arch/internal/pb"
	"time"
)

func OrderToProtobuf(order *database.Order) *pb.Order {
	items := make([]*pb.OrderItem, len(order.Items))
	for i, item := range order.Items {
		items[i] = OrderItemToProtobuf(&item)
	}
	return &pb.Order{
		Id:       order.ID,
		Customer: order.Customer,
		Date:     order.Date.UTC().Format(time.UnixDate),
		Total:    order.Total,
		Items:    items,
	}
}

func OrderItemToProtobuf(item *database.OrderItem) *pb.OrderItem {
	return &pb.OrderItem{
		Id:       item.ID,
		Product:  item.Product,
		Price:    item.Price,
		Quantity: int32(item.Quantity),
	}
}
