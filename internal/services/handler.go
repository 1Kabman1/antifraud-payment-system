package services

import (
	"encoding/json"
	"fmt"
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/hashStorage"
	"log"
	"net/http"
	"os"
	"strconv"
)

type apiHandler struct {
	s        hashStorage.Storage
	errorLog *log.Logger
	infoLog  *log.Logger
}

// NewApiHandler - set storage and logs
func NewApiHandler() apiHandler {
	return apiHandler{
		s:        hashStorage.NewStorage(),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}
}

// GetAggregationRules - Get aggregation data
func (h *apiHandler) GetAggregationRules(w http.ResponseWriter, _ *http.Request) {
	if h.s.RulesLen() == 0 {
		http.Error(w, "The rules don't exist yet", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Status", "success")
	w.Header().Set("Content-Type", "application/json")

	for _, aRule := range h.s.Rules() {

		ruleJson, err := json.Marshal(aRule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.errorLog.Println(err)
			return
		}

		if _, err := w.Write(ruleJson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte("\n")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.errorLog.Println(err)
			return
		}

	}
}

// CreateAggregationRule - create aggregation rule
func (h *apiHandler) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {

	aRule := newRule()

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&aRule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err)
		return
	}

	if h.s.IsRule(aRule.Name) {
		w.Header().Set("Message", "rule already exists")
		w.Header().Set("Status", " error "+strconv.Itoa(http.StatusConflict))

	} else {
		id := h.s.RulesLen() + 1
		aRule.AggregationRuleId = id
		h.s.SetRule(aRule.Name, aRule)
		w.Header().Set("Message", "rule "+strconv.Itoa(id)+" created")
		w.Header().Set("Status", "success")

	}

}

// RegisterOperation - counts aggregated based on the rules
func (h *apiHandler) RegisterOperation(w http.ResponseWriter, r *http.Request) {
	if h.s.RulesLen() == 0 {
		http.Error(w, "The rules don't exist yet", http.StatusInternalServerError)
		return
	}

	mapPING := map[string]interface{}{}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&mapPING); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err.Error())
		return
	}

	aggregatesBy, err := prepareTheDataForHashing(h.s.Rules(), mapPING)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err)
		return
	}

	for nameRule, keyCounter := range calculateHash(aggregatesBy) {
		err, tempRule := h.s.Rule(nameRule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.errorLog.Println(err)
			return
		}
		fmt.Println(keyCounter)
		aRule := tempRule.(rule)

		if h.s.IsCounter(keyCounter) {

			_, tmpCounter := h.s.Counter(keyCounter)
			c := tmpCounter.(counter)

			if aRule.AggregateValue == "count" {
				c.Value += 1
				h.s.SetCounter(keyCounter, c)
			} else {
				c.Value += int(mapPING["amount"].(float64))
				h.s.SetCounter(keyCounter, c)
			}
		} else {
			aNewCounter := newCounter()
			if aRule.AggregateValue == "amount" {
				aNewCounter.Value = int(mapPING["amount"].(float64))
			} else {
				aNewCounter.Value++
			}
			idCounter := h.s.CounterLen() + 1
			aNewCounter.id = idCounter
			h.s.SetCounter(keyCounter, aNewCounter)
			h.s.AddToArchivist(aRule.AggregationRuleId, idCounter)
		}
	}

}
