package services

import (
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestPrepareTheDataForHashing(t *testing.T) {

	rule1 := hashStorage.Rule{
		AggregationRuleId: 0,
		Name:              "rule1",
		AggregateBy:       []string{"a", "b"},
		AggregateValue:    "count",
		ExpirationTime:    2,
		TimePeriod:        2,
	}
	rule2 := hashStorage.Rule{
		AggregationRuleId: 1,
		Name:              "rule2",
		AggregateBy:       []string{"c", "d"},
		AggregateValue:    "amount",
		ExpirationTime:    2,
		TimePeriod:        2,
	}
	rules := map[int]*hashStorage.Rule{1: &rule1, 2: &rule2}
	payments := map[string]interface{}{"a": 1234.00, "b": "2", "c": "1234", "d": "2"}
	actual, _ := prepareTheDataForHashing(rules, payments)
	expected := map[int]string{0: "1234a2b0", 1: "1234c2d1"}
	assert.Equal(t, expected, actual)
}

func TestStorageIncreaseValue(t *testing.T) {
	var (
		expirationTime = 1
		timePeriod     = 2
	)
	rule1 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "count",
		ExpirationTime: expirationTime,
		TimePeriod:     timePeriod,
	}
	rule2 := hashStorage.Rule{
		Name:           "rule2",
		AggregateBy:    []string{"c", "d"},
		AggregateValue: "amount",
		ExpirationTime: expirationTime,
		TimePeriod:     timePeriod,
	}
	h := NewApiHandler()
	h.s.SetRule("rule1", &rule1)
	h.s.SetRule("rule2", &rule2)
	w := myResponseWriterTwo{}
	body := `{"a": 1, "b": "2", "c": "1", "d": "2", "amount": 100.00}`
	var expectedKeyCounter1 = [16]byte{144, 205, 130, 169, 81, 153, 100, 226, 115, 20, 161, 23, 204, 35, 219, 186}
	var expectedKeyCounter2 = [16]byte{120, 246, 85, 64, 24, 36, 1, 115, 223, 251, 47, 194, 246, 149, 48, 110}
	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))
	h.RegisterOperation(&w, r)
	_, c := h.s.Counter(expectedKeyCounter1)
	if ok := assert.EqualValues(t, c.LenTimeSeries(), expirationTime*60/timePeriod); !ok {
		panic(t.Error)
	}
	r, _ = http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))
	h.RegisterOperation(&w, r)
	_, c1 := h.s.Counter(expectedKeyCounter2)
	if ok := assert.EqualValues(t, c1.SumActual(), 200); !ok {
		panic(t.Error)
	}
	<-time.After(1 * time.Minute)
}

func TestTimerOfCounter(t *testing.T) {
	var (
		expirationTime = 0
		timePeriod     = 2
	)
	rule1 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "count",
		ExpirationTime: expirationTime,
		TimePeriod:     timePeriod,
	}
	rule2 := hashStorage.Rule{
		Name:           "rule2",
		AggregateBy:    []string{"c", "d"},
		AggregateValue: "amount",
		ExpirationTime: expirationTime,
		TimePeriod:     timePeriod,
	}
	h := NewApiHandler()
	h.s.SetRule("rule1", &rule1)
	h.s.SetRule("rule2", &rule2)
	w := myResponseWriterTwo{}
	body := `{"a": 1, "b": "2", "c": "1", "d": "2", "amount": 100.00}`
	var expectedKeyCounter2 = [16]byte{120, 246, 85, 64, 24, 36, 1, 115, 223, 251, 47, 194, 246, 149, 48, 110}
	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))
	h.RegisterOperation(&w, r)
	actual := expirationTime * 60 / timePeriod
	expected := 0
	if ok := assert.EqualValues(t, expected, actual); !ok {
		panic(t.Error)
	}
	r, _ = http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))
	h.RegisterOperation(&w, r)
	_, c1 := h.s.Counter(expectedKeyCounter2)
	if ok := assert.EqualValues(t, c1.SumActual(), 200); !ok {
		panic(t.Error)
	}
	<-time.After(1 * time.Minute)
}
