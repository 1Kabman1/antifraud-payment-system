package services

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Storage struct {
	idStatic int
	rules    map[string]*rule
	aCounter map[[16]byte]*counter
}

func NewStorage() *Storage {
	return &Storage{
		rules:    make(map[string]*rule),
		aCounter: make(map[[16]byte]*counter),
	}
}

func (s *Storage) GetAggregationData(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Status", "success")
	w.Header().Set("Content-Type", "application/json")

	for _, rule := range s.rules {

		jSON, err := json.Marshal(rule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err1 := w.Write(jSON)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}

		_, err1 = w.Write([]byte("\n"))
		if err1 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *Storage) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {

	newRule := new(rule)

	if r.Body != nil {
		defer r.Body.Close()
	}

	err := json.NewDecoder(r.Body).Decode(&newRule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, ok := s.rules[newRule.Name]

	if ok {
		w.Header().Set("Message", "rule already exists")
		w.Header().Set("Status", " error "+strconv.Itoa(http.StatusConflict))

	} else {
		s.idStatic++
		newRule.AggregationRuleId = s.idStatic
		s.rules[newRule.Name] = newRule
		w.Header().Set("Message", "rule "+strconv.Itoa(s.idStatic)+" created")
		w.Header().Set("Status", "success")
	}

}

func (s *Storage) CalculateTheAggregated(w http.ResponseWriter, r *http.Request) {
	mapPING := map[string]interface{}{}
	var amount float64
	if r.Body != nil {
		defer r.Body.Close()
	}

	err := json.NewDecoder(r.Body).Decode(&mapPING)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bul strings.Builder

	for _, ru := range s.rules {

		for _, agg := range ru.AggregateBy {
			if v, ok := mapPING[agg]; ok {
				switch inter := v.(type) {
				case float64:
					bul.WriteString(strconv.FormatFloat(inter, 'E', -1, 6))
					if agg == "amount" {
						amount = inter
					}
				case string:
					bul.WriteString(inter)
				}
			}
		}

		key := MD5String(bul.String())

		if c, ok := s.aCounter[key]; ok {
			if ru.AggregateValue == "count" {
				c.count++
			} else {
				c.amount += amount
			}
		} else {
			aCounter := newCounter()
			if aAmount, ok := mapPING["amount"].(float64); ok {
				aCounter.amount = aAmount
			} else {
				aCounter.count++
			}
			aCounter.summandId = s.idStatic
			s.aCounter[key] = aCounter
		}
		bul.Reset()
	}
}
