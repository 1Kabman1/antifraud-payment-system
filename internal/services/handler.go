package services

import (
	"encoding/json"
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

// CreateAggregationRule - create aggregation Rule
func (h *apiHandler) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {

	aRule := hashStorage.NewRule()

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&aRule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err)
		return
	}

	if h.s.HasRule(aRule.Name) {
		http.Error(w, "< Rule already exists \n", http.StatusConflict)
		return
	}

	h.s.SetRule(aRule.Name, &aRule)
	w.Write([]byte("Message " + "Rule " + strconv.Itoa(h.s.IdRule(aRule.Name)) + " created"))

}

// RegisterOperation - counts aggregated based on the rules
func (h *apiHandler) RegisterOperation(w http.ResponseWriter, r *http.Request) {

	payment := map[string]interface{}{}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err.Error())
		return
	}

	aggregatesBy, err := prepareTheDataForHashing(h.s.Rules(), payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err)
		return
	}

	for nameRule, keyCounter := range calculateHash(aggregatesBy) {
		err, aRule := h.s.Rule(nameRule)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.errorLog.Println(err)
			return
		}

		if !h.s.HasCounter(keyCounter) {
			h.s.SetCounter(keyCounter, aRule.Name)
		}

		_, c := h.s.Counter(keyCounter)

		if aRule.AggregateValue == count {
			c.Value += 1
		} else {
			c.Value += int(payment[amount].(float64))
		}
	}
}
