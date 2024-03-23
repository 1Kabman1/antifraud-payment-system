package services

import (
	"bytes"
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

type ApiHandler struct {
	s        hashStorage.Storage
	errorLog *log.Logger
	infoLog  *log.Logger
}

// NewApiHandler - set storage and logs
func NewApiHandler() ApiHandler {
	return ApiHandler{
		s:        hashStorage.NewStorage(),
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}
}

// GetAggregationRules - Get aggregation data
func (h *ApiHandler) GetAggregationRules(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]hashStorage.Rule, h.s.RulesLen())

	for _, aRule := range h.s.Rules() {
		resp[strconv.Itoa(aRule.AggregationRuleId)] = *aRule
	}

	ruleJson, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.errorLog.Println(err)
		return
	}
	if json.Valid(ruleJson) {
		if _, err = w.Write([]byte("Status " + "success \n")); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err = w.Write(ruleJson); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// CreateAggregationRule - create aggregation Rule
func (h *ApiHandler) CreateAggregationRule(w http.ResponseWriter, r *http.Request) {

	var buf bytes.Buffer
	aRule := hashStorage.NewRule()

	if r.Body != nil {
		defer func() {
			if err := r.Body.Close(); err != nil {
				h.errorLog.Println(err)
			}
		}()
	}

	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf.Bytes(), &aRule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.s.SetRule(aRule.Name, &aRule)

	if _, err := w.Write([]byte("Message " + "Rule " + strconv.Itoa(aRule.AggregationRuleId) + " created")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RegisterOperation - counts aggregated based on the rules
func (h *ApiHandler) RegisterOperation(w http.ResponseWriter, r *http.Request) {

	payment := map[string]interface{}{}

	if r.Body != nil {
		defer func() {
			if err := r.Body.Close(); err != nil {
				h.errorLog.Println(err)
			}
		}()
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
		err_, aRule := h.s.Rule(nameRule)
		if err_ != nil {
			http.Error(w, err_.Error(), http.StatusInternalServerError)
			h.errorLog.Println(err_)
			return
		}

		if !h.s.HasCounter(keyCounter) {
			h.s.SetCounter(keyCounter, aRule.AggregationRuleId)
		}

		_, c := h.s.Counter(keyCounter)

		if aRule.AggregateValue == count {
			c.Value += 1
		} else {
			c.Value += int(payment[amount].(float64))
		}
	}
}
