package services

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"strings"
	"testing"
)

type myResponseWriterTwo struct {
	t    testing.T
	flag bool
}

func (m *myResponseWriterTwo) Header() http.Header {
	return http.Header{}
}
func (m *myResponseWriterTwo) Write([]byte) (int, error) {

	return 0, nil
}
func (m *myResponseWriterTwo) WriteHeader(statusCode int) {}

func TestCalculateTheAggregated(t *testing.T) {

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

	h := NewApiHandler()
	h.s.SetRule("rule1", rule1)
	h.s.SetRule("rule2", rule2)

	w := myResponseWriterTwo{}

	body := `{"a": 1, "b": "2", "c": "1", "d": "2", "amount": 100.00}`

	var expectedKeyCounter1 = [16]byte{24, 159, 252, 74, 24, 132, 173, 22, 84, 162, 195, 1, 1, 133, 146, 20}
	var expectedKeyCounter2 = [16]byte{75, 47, 51, 201, 178, 229, 255, 233, 21, 199, 102, 59, 131, 78, 63, 209}

	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)

	_, counter1 := h.s.Counter(expectedKeyCounter1)
	c1 := counter1.(counter)
	if !assert.Equal(t, c1.Value, 1) {
		log.Panic()
	}

	r, _ = http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)

	_, counter2 := h.s.Counter(expectedKeyCounter2)
	c2 := counter2.(counter)
	if !assert.Equal(t, c2.Value, 200) {
		log.Panic()
	}

}
