package hashStorage

type Counter struct {
	id         int
	TotalValue int
	Values     *Node
}

func NewCounter() Counter {
	return Counter{}
}
