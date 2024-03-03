package hashStorage

type Rule struct {
	aggregationRuleId int
	Name              string   `json:"Name"`
	AggregateBy       []string `json:"AggregateBy"`
	AggregateValue    string   `json:"AggregateValue"` // enum aggregateType
}

func NewRule() Rule {
	return Rule{}
}
