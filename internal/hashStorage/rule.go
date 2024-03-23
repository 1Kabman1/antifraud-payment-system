package hashStorage

type Rule struct {
	AggregationRuleId int
	Name              string        `json:"Name"`
	AggregateBy       []string      `json:"AggregateBy"`
	AggregateValue    string        `json:"AggregateValue"`
	Duration          aTimeDuration `json:"Duration"`
	TimeFrom          aTime         `json:"TimeFrom"`
	TimeTo            aTime         `json:"TimeTo"`
}

func NewRule() Rule {
	return Rule{}
}
