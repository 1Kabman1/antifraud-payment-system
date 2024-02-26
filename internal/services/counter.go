package services

type counter struct {
	id    int
	Value int
}

func newCounter() counter {
	return counter{}
}
