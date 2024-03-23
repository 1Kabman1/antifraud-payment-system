package hashStorage

import "time"

type Counter struct {
	id         int
	Value      Order
	TotalValue int
}

func NewCounter() Counter {
	return Counter{}
}

type Order struct {
	Value int
	Time  time.Time
}
