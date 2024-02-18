package services

type counter struct {
	id     int
	count  int
	amount float64
}

func newCounter() counter {
	return counter{}
}
