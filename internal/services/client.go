package services

type Client struct {
	Id     int
	Name   string
	Amount int
	Count  int
}

type ListOfClients struct {
	Client map[int]*Client
}

func NewListOfClients() *ListOfClients {
	return &ListOfClients{
		Client: make(map[int]*Client),
	}
}
