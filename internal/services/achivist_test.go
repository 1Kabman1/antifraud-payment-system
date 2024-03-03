package services

import (
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"strings"
	"testing"
)

func Test_Archivist(t *testing.T) {

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
	h.s.SetRule("rule1", rule1)
	h.s.SetRule("rule2", rule2)

	w := myResponseWriterTwo{}

	body := `{"a": 1, "b": "2", "c": "1", "d": "2", "amount": 100.00}`

	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)

	body = `{"a": 12, "b": "23", "c": "1", "d": "42", "amount": 100.00}`

	r, _ = http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)

	count1, _ := h.s.Archivist("1")

	if !assert.Equal(t, count1, 2) {
		log.Panic()
	}

	count2, _ := h.s.Archivist("2")

	if !assert.Equal(t, count2, 2) {
		log.Panic()
	}

}
