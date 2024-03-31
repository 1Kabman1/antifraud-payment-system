package services

import (
	"encoding/json"
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

type myResponseWriter struct {
	t    testing.T
	flag bool
}

func (m *myResponseWriter) Header() http.Header {
	return http.Header{}
}
func (m *myResponseWriter) Write(w []byte) (int, error) {
	if m.flag {
		aRule := hashStorage.Rule{
			AggregationRuleId: 1,
			Name:              "rule1",
			AggregateBy:       []string{"a", "b"},
			AggregateValue:    "count",
		}

		resp := make(map[string]hashStorage.Rule)

		resp[strconv.Itoa(aRule.AggregationRuleId)] = aRule

		ruleJson, _ := json.Marshal(resp)

		ok := assert.Equal(&m.t, string(w), string(ruleJson))
		if !ok {
			m.t.Fatal(w)
		}
	}
	m.flag = true
	return 0, nil
}
func (m *myResponseWriter) WriteHeader(_ int) {}

func TestAggregationData(t *testing.T) {

	rule1 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "count",
	}

	h := NewApiHandler()
	h.s.SetRule("rule1", &rule1)

	r := http.Request{}
	m := myResponseWriter{}
	h.GetAggregationRules(&m, &r)

}
