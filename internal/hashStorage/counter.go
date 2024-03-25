package hashStorage

import "time"

type Counter struct {
	id         int
	Value      []Order
	TotalValue int
}

func NewCounter() Counter {
	return Counter{Value: make([]Order, 0)}
}

type Order struct {
	Value int
	Time  time.Time
}

func NewOrder() Order {
	return Order{}
}
