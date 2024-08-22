package entity

type OrderRepository interface {
	Save(order *Order) error
	List(limit, offset int32) ([]Order, error)
}
