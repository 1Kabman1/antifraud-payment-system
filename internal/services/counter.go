package services

type counter struct {
	summandId int
	count     int
	amount    float64
}

func newCounter() *counter {
	return &counter{}
}
