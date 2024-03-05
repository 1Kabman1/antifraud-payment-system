package hashStorage

type Counter struct {
	id    int
	Value int
}

func NewCounter() Counter {
	return Counter{}
}
