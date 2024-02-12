package services

type Rule struct {
	AggregationRuleId int      `json:"id"`
	AggregateBy       []string `json:"AggregateBy"`
}

type RuleBook struct {
	Id    int
	Rules map[[16]byte]*Rule
}

func NewRuleBook() *RuleBook {
	return &RuleBook{
		Rules: make(map[[16]byte]*Rule),
	}
}
