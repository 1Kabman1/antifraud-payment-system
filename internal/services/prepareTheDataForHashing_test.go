package services

import (
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrepareTheDataForHashing(t *testing.T) {

	rule1 := hashStorage.Rule{
		AggregationRuleId: 0,
		Name:              "rule1",
		AggregateBy:       []string{"a", "b"},
		AggregateValue:    "count",
	}
	rule2 := hashStorage.Rule{
		AggregationRuleId: 1,
		Name:              "rule2",
		AggregateBy:       []string{"c", "d"},
		AggregateValue:    "amount",
	}

	rules := map[int]*hashStorage.Rule{1: &rule1, 2: &rule2}

	payments := map[string]interface{}{"a": 1234.00, "b": "2", "c": "1234", "d": "2"}

	actual, _ := prepareTheDataForHashing(rules, payments)

	expected := map[int]string{0: "1234a2b0", 1: "1234c2d1"}
	assert.Equal(t, expected, actual)
}
