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

	jDataCh := make(chan []byte)

	go func() {
		for _, obj := range s.mp {

			j, err := json.Marshal(obj)
			if err != nil {
				close(jDataCh)
			}
			jDataCh <- j

		}
		close(jDataCh)
	}()

	for {
		rules, ok := <-jDataCh
		if ok {
			w.Write(rules)
			w.Write([]byte(("\n")))
		} else {
			break
		}
	}

}

func (s *Storage) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {

	newRule := new(serv.Rule)

	defer r.Body.Close()
	agrByChIn := make(chan []string)
	argByChOut := make(chan [16]byte)
	idOut := ""

	go func() {
		err := json.NewDecoder(r.Body).Decode(&newRule)
		if err != nil {
			w.Header().Add("Status", "unsuccessful")
			return
		}
		newRule.AggregateBy = append(newRule.AggregateBy, newRule.Name)
		agrByChIn <- newRule.AggregateBy
	}()

	go serv.MD5(<-agrByChIn, argByChOut)

	argBy := <-argByChOut

	_, ok := s.mp[argBy]

	if ok {
		s.mp[argBy].Amount += newRule.Amount
		s.mp[argBy].Count += 1
		idOut = " exists"
	} else {
		s.idStatic++
		newRule.AggregationRuleId = s.idStatic
		s.mp[argBy] = newRule
		s.mp[argBy].Count = 1
		idOut = strconv.Itoa(s.idStatic)
		idOut += " created"
	}
	w.Header().Set("Message", "Rule "+idOut)
	w.Header().Set("Status", "success")

}
