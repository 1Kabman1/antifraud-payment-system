package rest

import (
	"encoding/json"
	serv "github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"net/http"
	"strconv"
)

func (s *Storage) GetAggregationData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Status", "success")
	w.Header().Set("Content-Type", "appliction/json")
	jDataCh := make(chan []byte)

	go func() {
		for _, obj := range s.mp {
			go func() {
				j, err := json.Marshal(obj)
				if err != nil {
					return
				}
				jDataCh <- j
			}()

		}
	}()

	for i := 0; i < len(s.mp); i++ {
		select {
		case rulls := <-jDataCh:
			w.Write(rulls)
			w.Write([]byte(("\n")))
		}

	}

}

func (s *Storage) CreateAggregationRoull(w http.ResponseWriter, r *http.Request) {

	newRull := new(serv.Rull)
	defer r.Body.Close()
	agrByChIn := make(chan []string)
	argByChOut := make(chan [16]byte)
	idOut := ""

	go func() {
		err := json.NewDecoder(r.Body).Decode(&newRull)
		if err != nil {
			w.Header().Add("Status", "unsuccessful")
			return
		}
		agrByChIn <- newRull.AggregateBy
	}()

	go serv.MD5(<-agrByChIn, argByChOut)

	argBy := <-argByChOut

	_, ok := s.mp[argBy]

	if ok {
		s.mp[argBy].Amount += newRull.Amount
		s.mp[argBy].Count += 1
		idOut = " exists"
	} else {
		s.idStatic++
		newRull.AggregationRuleId = s.idStatic
		s.mp[argBy] = newRull
		s.mp[argBy].Count = 1
		idOut = strconv.Itoa(s.idStatic)
		idOut += "created"
	}
	w.Header().Set("Message", "Rule "+idOut)
	w.Header().Set("Status", "success")

}
