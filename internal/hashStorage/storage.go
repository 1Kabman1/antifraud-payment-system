package hashStorage

import (
	"errors"
	"strconv"
)

type Storage struct {
	rules     map[string]Rule
	counter   map[[16]byte]interface{}
	archivist map[int][]int
}

// NewStorage - create a Storage
func NewStorage() Storage {
	return Storage{
		rules:     make(map[string]Rule),
		counter:   make(map[[16]byte]interface{}),
		archivist: make(map[int][]int, 5),
	}
}

// SetRule - set rule in map
func (s *Storage) SetRule(nameKey string, rule Rule) {
	id := s.RulesLen() + 1
	rule.aggregationRuleId = id
	s.rules[nameKey] = rule
}

// RulesLen - returns the length of the map
func (s *Storage) RulesLen() int {
	return len(s.rules)
}

// IdRule - Return id rule
func (s *Storage) IdRule(name string) int {
	return s.rules[name].aggregationRuleId
}

// Rules - returns rules
func (s *Storage) Rules() map[string]Rule {
	return s.rules
}

// Rule - returns a rule
func (s *Storage) Rule(key string) (error, Rule) {
	err := errors.New("Key is not correct")

	_, ok := s.rules[key]
	if ok {
		return nil, s.rules[key]
	}

	return err, Rule{}
}

// HasRule - return bool
func (s *Storage) HasRule(key string) bool {
	_, ok := s.rules[key]
	return ok
}

// HasCounter - Checks if there is a value in the map
func (s *Storage) HasCounter(key [16]byte) bool {
	_, ok := s.counter[key]

	return ok

}

// SetCounter - sets id for c
func (s *Storage) SetCounter(key [16]byte, v interface{}) {
	s.counter[key] = v

}

// Counter - return counter
func (s *Storage) Counter(key [16]byte) (error, interface{}) {
	err := errors.New("Key is not correct")

	_, ok := s.counter[key]
	if ok {
		return nil, s.counter[key]
	}
	return err, nil
}

// CounterLen - return len
func (s *Storage) CounterLen() int {
	return len(s.counter)
}

// AddToArchivist - filled in by the archivist
func (s *Storage) AddToArchivist(idRule, idCounter int) {
	s.archivist[idRule] = append(s.archivist[idRule], idCounter)
}

// Archivist - is for test
func (s *Storage) Archivist(key string) (int, error) {
	id, err := strconv.Atoi(key)
	if err != nil {
		return 0, err
	}
	return len(s.archivist[id]), nil
}
