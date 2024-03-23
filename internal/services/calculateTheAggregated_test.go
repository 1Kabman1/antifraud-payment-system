package services

import (
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
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

	rule1 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "count",
	}
	rule2 := hashStorage.Rule{
		Name:           "rule2",
		AggregateBy:    []string{"c", "d"},
		AggregateValue: "amount",
	}

	h := NewApiHandler()
	h.s.SetRule("rule1", &rule1)
	h.s.SetRule("rule2", &rule2)

	w := myResponseWriterTwo{}

	body := `{"a": 1, "b": "2", "c": "1", "d": "2", "amount": 100.00}`

	var expectedKeyCounter1 = [16]byte{160, 160, 128, 244, 46, 111, 19, 179, 162, 223, 19, 63, 7, 48, 149, 221}
	var expectedKeyCounter2 = [16]byte{76, 86, 255, 76, 228, 170, 249, 87, 58, 165, 223, 249, 19, 223, 153, 122}

	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)

	_, counter1 := h.s.Counter(expectedKeyCounter1)
	if !assert.Equal(t, counter1.Value.Value, 100) {
		log.Panic()
	}

	r, _ = http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)

	_, counter2 := h.s.Counter(expectedKeyCounter2)
	if !assert.Equal(t, counter2.Value.Value, 2) {
		log.Panic()
	}

}

func TestCalculateTheAggregatedIdenticalAggregateBy(t *testing.T) {

	rule1 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "amount",
	}
	rule2 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "amount",
	}

	h := NewApiHandler()
	h.s.SetRule("rule1", &rule1)
	h.s.SetRule("rule1", &rule2)

	w := myResponseWriterTwo{}

	body := `{"a": 1, "b": "2", "c": "1", "d": "2", "amount": 100.00}`

	var expectedKeyCounter1 = [16]byte{144, 205, 130, 169, 81, 153, 100, 226, 115, 20, 161, 23, 204, 35, 219, 186}

	var expectedKeyCounter2 = [16]byte{160, 195, 89, 108, 8, 85, 162, 177, 34, 120, 114, 137, 155, 21, 128, 12}

	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)

	_, counter1 := h.s.Counter(expectedKeyCounter1)
	_, counter2 := h.s.Counter(expectedKeyCounter2)
	if !assert.Equal(t, counter1.Value, counter2.Value) {
		log.Panic()
	}

}
