package hashStorage

type Rule struct {
	AggregationRuleId int      `json:"AggregationRuleId"`
	Name              string   `json:"Name"`
	AggregateBy       []string `json:"AggregateBy"`
	AggregateValue    string   `json:"AggregateValue"` // enum aggregateType
}

func NewRule() Rule {
	return Rule{}
}
