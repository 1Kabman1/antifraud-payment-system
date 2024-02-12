package services

type Client struct {
	Name   string
	Amount int
	Count  int
}

type ListOfClients struct {
	Client map[string]*Client
}

func NewListOfClients() *ListOfClients {
	return &ListOfClients{
		Client: make(map[string]*Client),
	}
}
