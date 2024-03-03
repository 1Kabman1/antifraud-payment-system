package services

import (
	"encoding/json"
	"fmt"
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	"github.com/stretchr/testify/assert"
	"net/http"
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
	if !m.flag {
		aRule := hashStorage.Rule{
			Name:           "rule1",
			AggregateBy:    []string{"a", "b"},
			AggregateValue: "count",
		}

		fmt.Println(w)

		expected, _ := json.Marshal(aRule)

		ok := assert.Equal(&m.t, string(w), string(expected))
		if !ok {
			m.t.Fatal(w)
		}
	}
	m.flag = true
	return 0, nil
}
func (m *myResponseWriter) WriteHeader(statusCode int) {}

func TestAggregationData(t *testing.T) {

	rule1 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "count",
	}

	h := NewApiHandler()
	h.s.SetRule("rule1", rule1)

	r := http.Request{}
	m := myResponseWriter{}
	h.GetAggregationRules(&m, &r)

}
