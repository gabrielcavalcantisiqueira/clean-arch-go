package entity

type OrderRepositoryInterface interface {
	GetOrders() ([]Order, error)
	Save(order *Order) error
	// GetTotal() (int, error)
}
