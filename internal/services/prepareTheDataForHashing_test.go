package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareTheDataForHashing(t *testing.T) {

	rule1 := rule{
		AggregationRuleId: 1,
		Name:              "rule1",
		AggregateBy:       []string{"a", "b"},
		AggregateValue:    "count",
	}
	rule2 := rule{
		AggregationRuleId: 2,
		Name:              "rule2",
		AggregateBy:       []string{"c", "d"},
		AggregateValue:    "amount",
	}

	rules := map[string]interface{}{}
	rules["rule1"] = rule1
	rules["rules2"] = rule2

	payments := map[string]interface{}{}
	payments["a"] = 1234.00
	payments["b"] = "2"
	payments["c"] = "1234"
	payments["d"] = "2"

	actual, _ := prepareTheDataForHashing(rules, payments)

	expected := map[string]string{"rule1": "1234a2b", "rule2": "1234c2d"}
	assert.Equal(t, expected, actual)
}
