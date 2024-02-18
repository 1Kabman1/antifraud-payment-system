package services

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func (s *Storage) GetAggregationData(w http.ResponseWriter, _ *http.Request) {
	if len(s.rules) == 0 {
		http.Error(w, "The rules don't exist yet", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Status", "success")
	w.Header().Set("Content-Type", "application/json")

	for _, aRule := range s.rules {

		ruleJson, err := json.Marshal(aRule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write(ruleJson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte("\n")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *Storage) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {

	aRule := newRule()

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&aRule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, ok := s.rules[aRule.Name]

	if ok {
		w.Header().Set("Message", "rule already exists")
		w.Header().Set("Status", " error "+strconv.Itoa(http.StatusConflict))

	} else {
		s.idRules++
		aRule.AggregationRuleId = s.idRules
		s.rules[aRule.Name] = aRule
		w.Header().Set("Message", "rule "+strconv.Itoa(s.idRules)+" created")
		w.Header().Set("Status", "success")
	}

}

func (s *Storage) CalculateTheAggregated(w http.ResponseWriter, r *http.Request) {
	if len(s.rules) == 0 {
		http.Error(w, "The rules don't exist yet", http.StatusInternalServerError)
		return
	}

	var aBuilder strings.Builder
	mapPING := map[string]interface{}{}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&mapPING); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ws := sync.WaitGroup{}
	chan1 := make(chan map[string]string)
	chan2 := make(chan map[string][16]byte)

	go func() {

		aggregatesBy := make(map[string]string, len(s.rules))

		for key, aRule := range s.rules {

			for _, agg := range aRule.AggregateBy {
				if v, ok := mapPING[agg]; ok {
					switch aInterface := v.(type) {
					case float64:
						aBuilder.WriteString(strconv.FormatFloat(aInterface, 'E', -1, 64))
					case string:
						aBuilder.WriteString(aInterface)
					}
					aggregate := aBuilder.String()
					aggregatesBy[key] = aggregate
					aBuilder.Reset()
				}
			}
		}
		chan1 <- aggregatesBy
	}()

	go func() {
		MD5(chan1, chan2)
	}()

	ws.Add(1)
	go func() {
		defer ws.Done()

		for nameRule, keyCounter := range <-chan2 {

			if c, ok := s.counters[keyCounter]; ok {
				if s.rules[nameRule].AggregateValue == "count" {
					c.count++
				} else {
					c.amount += mapPING["amount"].(float64)
				}
			} else {
				aNewCounter := newCounter()
				if s.rules[nameRule].AggregateValue == "amount" {
					aNewCounter.amount = mapPING["amount"].(float64)
				} else {
					aNewCounter.count++
				}
				s.idCounters++
				aNewCounter.id = s.idCounters
				s.counters[keyCounter] = aNewCounter
			}
		}
	}()

	ws.Wait()
	close(chan2)
	close(chan1)
}
