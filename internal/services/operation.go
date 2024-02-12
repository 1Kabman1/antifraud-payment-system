package services

type Operation struct {
	Name        string   `json:"Name"`
	AggregateBy []string `json:"AggregateBy"`
	Amount      int      `json:"AggregatedValue"`
}
