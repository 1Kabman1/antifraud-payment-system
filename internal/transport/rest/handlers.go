package rest

import (
	"encoding/json"
	serv "github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"net/http"
	"strconv"
	"sync"
)

func (s *Storage) GetAggregationData(w http.ResponseWriter, _ *http.Request) {
	//
	//w.Header().Set("Status", "success")
	//w.Header().Set("Content-Type", "application/json")
	//
	//jDataCh := make(chan []byte)
	//
	//go func() {
	//	for _, obj := range s.mp {
	//
	//		j, err := json.Marshal(obj)
	//		if err != nil {
	//			close(jDataCh)
	//		}
	//		jDataCh <- j
	//
	//	}
	//	close(jDataCh)
	//}()
	//
	//for {
	//	rules, ok := <-jDataCh
	//	if ok {
	//		w.Write(rules)
	//		w.Write([]byte(("\n")))
	//	} else {
	//		break
	//	}
	//}

}

func (s *Storage) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup

	newOperation := new(serv.Operation)
	newRule := new(serv.Rule)
	newClient := new(serv.Client)

	defer r.Body.Close()
	aggregateByCHAN := make(chan []string)
	ruleCHAN := make(chan [16]byte)
	operationCHAN := make(chan serv.Operation)
	idOut := ""

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := json.NewDecoder(r.Body).Decode(&newOperation)
		if err != nil {
			w.Header().Add("Status", "unsuccessful")
			return
		}

		aggregateByCHAN <- newOperation.AggregateBy
		operationCHAN <- *newOperation
	}()

	go func() {
		defer close(aggregateByCHAN)
		serv.MD5(<-aggregateByCHAN, ruleCHAN)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rule := <-ruleCHAN
		s.mx.Lock()
		defer s.mx.Unlock()
		_, ok := s.ruleBook.Rules[rule]
		if ok {
			s.ruleBook.Rules[rule].AggregateBy = newOperation.AggregateBy
			idOut = strconv.Itoa(s.ruleBook.Rules[rule].AggregationRuleId)
		} else {
			newId := s.GiveIdForRules()
			newRule.AggregationRuleId = newId
			newRule.AggregateBy = newOperation.AggregateBy
			s.ruleBook.Rules[rule] = newRule
			idOut = strconv.Itoa(newId)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		operation := <-operationCHAN
		s.mx.Lock()
		defer s.mx.Unlock()
		_, ok := s.listClients.Client[operation.Name]
		if ok {
			s.listClients.Client[operation.Name].Count += 1
			s.listClients.Client[operation.Name].Amount += operation.Amount
		} else {
			newClient.Name = operation.Name
			newClient.Amount = operation.Amount
			newClient.Count += 1
			s.listClients.Client[operation.Name] = newClient
		}
	}()

	wg.Wait()
	w.Header().Set("Message", "Rule "+idOut)
	w.Header().Set("Status", "success")
}
