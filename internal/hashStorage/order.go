package hashStorage

type Order struct {
	value int
	t     aTimeDuration
}

func NewOrder() Order {
	return Order{}

}
