package rest

import (
	"encoding/json"
	serv "github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"net/http"
	"strconv"
	"sync"
)

func (s *Storage) GetAggregationData(w http.ResponseWriter, _ *http.Request) {

	rulesCHAN := make(chan []byte)
	var ws sync.WaitGroup

	ws.Add(1)
	go func() {
		defer ws.Done()
		for _, rule := range s.rules {
			ws.Add(1)
			go func(rule interface{}) {
				defer ws.Done()
				jSON, err := json.Marshal(rule)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				rulesCHAN <- jSON
			}(rule)
		}
	}()

	go func() {
		ws.Wait()
		close(rulesCHAN)
	}()

	w.Header().Set("Status", "success")
	w.Header().Set("Content-Type", "application/json")

	for rule := range rulesCHAN {
		_, err := w.Write(rule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write([]byte("\n"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	ws.Wait()
}

func (s *Storage) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {

	var ws sync.WaitGroup
	newRule := new(serv.Rule)
	agrByCHAN := make(chan []string)
	hashCHAN := make(chan [16]byte)

	defer r.Body.Close()

	ws.Add(1)
	go func() {
		defer ws.Done()
		err := json.NewDecoder(r.Body).Decode(&newRule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newRule.AggregateBy = append(newRule.AggregateBy, newRule.Name)
		agrByCHAN <- newRule.AggregateBy
	}()

	ws.Add(1)
	go func() {
		defer ws.Done()
		defer close(agrByCHAN)
		serv.MD5(<-agrByCHAN, hashCHAN)
	}()

	ws.Add(1)
	go func() {
		defer ws.Done()
		defer close(hashCHAN)

		key := <-hashCHAN

		_, ok := s.rules[key]

		if ok {
			w.Header().Set("Message", "rule already exists")
			w.Header().Set("Status", " error "+strconv.Itoa(http.StatusConflict))

		} else {
			s.idStatic++
			newRule.AggregationRuleId = s.idStatic
			s.rules[key] = newRule

			w.Header().Set("Message", "Rule "+strconv.Itoa(s.idStatic)+" created")
			w.Header().Set("Status", "success")
		}
	}()

	ws.Wait()

}
