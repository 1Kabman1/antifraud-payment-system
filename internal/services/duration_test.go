package services

import (
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	_ "github.com/stretchr/testify/assert"
	_ "log"
	"net/http"
	"strings"
	"testing"
	_ "time"
)

func TestStorage_IncreaseValue(t *testing.T) {

	rule1 := hashStorage.Rule{
		Name:           "rule1",
		AggregateBy:    []string{"a", "b"},
		AggregateValue: "count",
		Duration:       struct{ Duration int64 }{Duration: 100},
	}
	rule2 := hashStorage.Rule{
		Name:           "rule2",
		AggregateBy:    []string{"c", "d"},
		AggregateValue: "amount",
		Duration:       struct{ Duration int64 }{Duration: 1000},
	}

	h := NewApiHandler()
	h.s.SetRule("rule1", &rule1)
	h.s.SetRule("rule2", &rule2)

	w := myResponseWriterTwo{}

	body := `{"a": 1, "b": "2", "c": "1", "d": "2", "amount": 100.00}`

	var expectedKeyCounter1 = [16]byte{144, 205, 130, 169, 81, 153, 100, 226, 115, 20, 161, 23, 204, 35, 219, 186}
	//var expectedKeyCounter2 = [16]byte{120, 246, 85, 64, 24, 36, 1, 115, 223, 251, 47, 194, 246, 149, 48, 110}

	r, _ := http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	h.RegisterOperation(&w, r)
	_, c := h.s.Counter(expectedKeyCounter1)

	value := (c.Values.Front()).Value.(hashStorage.Order)

	//_, counter1 := h.s.Counter(expectedKeyCounter1)
	//if !assert.Equal(t, counter1.TotalValue, 1) {
	//	log.Panic()
	//}

	//r, _ = http.NewRequest("POST", "http://127.0.0.1:8080/register", strings.NewReader(body))

	//h.RegisterOperation(&w, r)

	//_, counter2 := h.s.Counter(expectedKeyCounter2)
	//if !assert.Equal(t, counter2.TotalValue, 200) {
	//	log.Panic()
	//}

}

func TestStorage_SetCounter(t *testing.T) {

}
