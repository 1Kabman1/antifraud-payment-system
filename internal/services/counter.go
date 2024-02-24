package services

type counter struct {
	id    int
	value int
}

func newCounter() counter {
	return counter{}
}
