package services

type rule struct {
	AggregationRuleId int      `json:"id"`
	Name              string   `json:"Name"`
	AggregateBy       []string `json:"AggregateBy"`
	AggregateValue    string   `json:"AggregateValue"` // enum aggregateType
}

func newRule() rule {
	return rule{}
}
