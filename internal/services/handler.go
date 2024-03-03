package services

import (
	"encoding/json"
	"fmt"
	"github.com/1Kabman1/antifraud-payment-system/internal/hashStorage"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	count  string = "count"
	amount        = "amount"
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

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Status", "success")

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

	//добавить хеш по агругирующим в правиле , есть баг имена у правил разные а агриг рующие одинаковые
	if h.s.HasRule(aRule.Name) {
		http.Error(w, "< Rule already exists \n", http.StatusConflict)
		return
	} else {
		id := h.s.RulesLen() + 1
		aRule.AggregationRuleId = id
		h.s.SetRule(aRule.Name, aRule)
		w.Header().Add("Message", "rule "+strconv.Itoa(id)+" created")
		w.Header().Add("Status", "success")

	}

}

// RegisterOperation - counts aggregated based on the rules
func (h *apiHandler) RegisterOperation(w http.ResponseWriter, r *http.Request) {

	mapping := map[string]interface{}{}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&mapping); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err.Error())
		return
	}

	aggregatesBy, err := prepareTheDataForHashing(h.s.Rules(), mapping)
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

		if !h.s.HasCounter(keyCounter) {
			aNewCounter := newCounter()
			idCounter := h.s.CounterLen() + 1
			aNewCounter.id = idCounter
			h.s.SetCounter(keyCounter, aNewCounter)
			h.s.AddToArchivist(aRule.AggregationRuleId, idCounter)
		}

		_, tmpCounter := h.s.Counter(keyCounter)
		c := tmpCounter.(counter)

		if aRule.AggregateValue == count {
			c.Value += 1
			h.s.SetCounter(keyCounter, c)
		} else {
			c.Value += int(mapping[amount].(float64))
			h.s.SetCounter(keyCounter, c)
		}
	}
}
