package hashStorage

type Order struct {
	Value int
	T     timeDuration
}

func NewOrder() Order {
	return Order{}

}
