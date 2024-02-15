package rest

import (
	"encoding/json"
	serv "github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"net/http"
	"strconv"
)

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

	newRule := new(serv.Rule)

	defer r.Body.Close()

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
		w.Header().Set("Message", "Rule "+strconv.Itoa(s.idStatic)+" created")
		w.Header().Set("Status", "success")
	}

}
