package hashStorage

type Rule struct {
	AggregationRuleId int
	Name              string       `json:"Name"`
	AggregateBy       []string     `json:"AggregateBy"`
	AggregateValue    string       `json:"AggregateValue"`
	Duration          timeDuration `json:"DurationSec"`
}

func NewRule() Rule {
	return Rule{}
}
