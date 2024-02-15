package services

type Counter struct {
	summandId int
	count     int
	amount    int
}

func NewCounter() *Counter {
	return &Counter{}
}
