package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Storage struct {
	idRules    int
	idCounters int
	rules      map[string]*rule
	counters   map[[16]byte]*counter
	mux        sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		rules:    make(map[string]*rule),
		counters: make(map[[16]byte]*counter),
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

		if _, err := w.Write(jSON); err != nil {
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

	newRule := new(rule)

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&newRule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, ok := s.rules[newRule.Name]

	if ok {
		w.Header().Set("Message", "rule already exists")
		w.Header().Set("Status", " error "+strconv.Itoa(http.StatusConflict))

	} else {
		s.idRules++
		newRule.AggregationRuleId = s.idRules
		s.rules[newRule.Name] = newRule
		w.Header().Set("Message", "rule "+strconv.Itoa(s.idRules)+" created")
		w.Header().Set("Status", "success")
	}

}

func (s *Storage) CalculateTheAggregated(w http.ResponseWriter, r *http.Request) {
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
	md5CHAN := make(chan []string)
	keyCHAN := make(chan [][16]byte)

	ws.Add(1)
	go func() {
		defer ws.Done()

		aggregateByS := make([]string, len(s.rules))

		for _, aRule := range s.rules {

			for _, agg := range aRule.AggregateBy {
				if v, ok := mapPING[agg]; ok {
					//s.mux.Lock()
					switch inter := v.(type) {
					case float64:
						aBuilder.WriteString(strconv.FormatFloat(inter, 'E', -1, 64))
						if agg == "amount" {
						}
					case string:
						aBuilder.WriteString(inter)
					}
					aAggregateBy := aBuilder.String()
					//s.mux.Unlock()

					aggregateByS = append(aggregateByS, aAggregateBy)
					aBuilder.Reset()
				}
			}
		}
		md5CHAN <- aggregateByS

	}()

	ws.Add(1)
	go func() {
		defer ws.Done()
		MD5(md5CHAN, keyCHAN)
	}()

	ws.Add(1)
	go func() {
		defer ws.Done()

		for _, key := range <-keyCHAN {

			s.mux.Lock()
			if c, ok := s.counters[key]; ok {
				if aRule.AggregateValue == "count" {
					c.count++
				} else {
					c.amount += mapPING["amount"].(float64)
				}
			} else {
				aCounter := newCounter()
				if aAmount, ok := mapPING["amount"].(float64); ok {
					aCounter.amount = aAmount
				} else {
					aCounter.count++
				}
				s.idCounters++
				aCounter.id = s.idCounters
				s.counters[key] = aCounter
			}
			s.mux.Unlock()

		}
	}()

	ws.Wait()
	close(keyCHAN)
	close(md5CHAN)

	fmt.Println(s.counters)
}

//func (s *Storage) CalculateTheAggregated(w http.ResponseWriter, r *http.Request) {
//
//	var bul strings.Builder
//	mapPING := map[string]interface{}{}
//
//	if r.Body != nil {
//		defer r.Body.Close()
//	}
//
//	if err := json.NewDecoder(r.Body).Decode(&mapPING); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	ws := sync.WaitGroup{}
//	md5CHAN := make(chan string)
//	keyCHAN := make(chan [16]byte)
//
//	ws.Add(1)
//	go func() {
//		defer ws.Done()
//
//		for _, aRule := range s.rules {
//
//			ws.Add(1)
//			go func() {
//				defer ws.Done()
//
//				for _, agg := range aRule.AggregateBy {
//
//					if v, ok := mapPING[agg]; ok {
//						switch inter := v.(type) {
//						case float64:
//							bul.WriteString(strconv.FormatFloat(inter, 'E', -1, 64))
//							if agg == "amount" {
//							}
//						case string:
//							bul.WriteString(inter)
//						}
//					}
//
//				}
//				//s.mux.Lock()
//				//aAggregateBy := bul.String()
//
//				//s.mux.Unlock()
//				md5CHAN <- bul.String()
//				bul.Reset()
//
//			}()
//			ws.Add(1)
//			go func() {
//				defer ws.Done()
//				MD5(md5CHAN, keyCHAN)
//			}()
//
//			ws.Add(1)
//			go func() {
//				defer ws.Done()
//				key := <-keyCHAN
//
//				//s.mux.Lock()
//				if c, ok := s.counters[key]; ok {
//					if aRule.AggregateValue == "count" {
//						c.count++
//					} else {
//						c.amount += mapPING["amount"].(float64)
//					}
//				} else {
//					aCounter := newCounter()
//					if aAmount, ok := mapPING["amount"].(float64); ok {
//						aCounter.amount = aAmount
//					} else {
//						aCounter.count++
//					}
//					aCounter.id = s.idRules
//					s.counters[key] = aCounter
//					fmt.Println(s.counters[key])
//
//				}
//				//s.mux.Unlock()
//
//			}()
//			ws.Wait()
//		}
//
//	}()
//	//ws.Wait()
//	close(keyCHAN)
//	close(md5CHAN)
//
//	fmt.Println(s.counters)
//}
