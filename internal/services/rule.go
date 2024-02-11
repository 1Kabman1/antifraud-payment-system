package services

type Rule struct {
	AggregationRuleId int      `json:"id"`
	Name              string   `json:"Name"`
	AggregateBy       []string `json:"AggregateBy"`
	Amount            int      `json:"AggregatedValue"`
	Count             int      `json:"Count"`
}
