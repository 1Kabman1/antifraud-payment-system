package hashStorage

import (
	"errors"
	"strconv"
)

type Storage struct {
	rules     map[int]*Rule
	counter   map[[16]byte]*Counter
	archivist map[int][]int
}

// NewStorage - create a Storage
func NewStorage() Storage {
	return Storage{
		rules:     make(map[int]*Rule),
		counter:   make(map[[16]byte]*Counter),
		archivist: make(map[int][]int, 5),
	}
}

// SetRule - set rule in map
func (s *Storage) SetRule(name string, rule *Rule) {
	id := s.RulesLen() + 1
	rule.AggregationRuleId = id
	rule.Name = name
	s.rules[id] = rule

}

// RulesLen - returns the length of the map
func (s *Storage) RulesLen() int {
	return len(s.rules)
}

// Rules - returns rules
func (s *Storage) Rules() map[int]*Rule {
	return s.rules
}

// Rule - returns a rule
func (s *Storage) Rule(id int) (error, *Rule) {
	_, ok := s.rules[id]
	if ok {
		r := s.rules[id]
		return nil, r
	}
	return errors.New("Key is not correct"), &Rule{}
}

// HasRule - return bool
func (s *Storage) HasRule(id int) bool {
	_, ok := s.rules[id]
	return ok
}

// HasCounter - Checks if there is a value in the map
func (s *Storage) HasCounter(key [16]byte) bool {
	_, ok := s.counter[key]
	return ok
}

// SetCounter - sets id for c
func (s *Storage) SetCounter(key [16]byte, idRule int) {
	aNewCounter := NewCounter()
	idCounter := s.CounterLen() + 1
	aNewCounter.id = idCounter
	s.counter[key] = &aNewCounter
	s.AddToArchivist(idRule, idCounter)
}

// Counter - return Counter
func (s *Storage) Counter(key [16]byte) (error, *Counter) {
	_, ok := s.counter[key]
	if ok {
		c := s.counter[key]
		return nil, c
	}
	return errors.New("Key is not correct"), &Counter{}
}

// CounterLen - return len
func (s *Storage) CounterLen() int {
	return len(s.counter)
}

// AddToArchivist - filled in by the archivist
func (s *Storage) AddToArchivist(idRule, idCounter int) {
	s.archivist[idRule] = append(s.archivist[idRule], idCounter)
}

// ArchivistForTest - is for test
func (s *Storage) ArchivistForTest(key string) (int, error) {
	id, err := strconv.Atoi(key)
	if err != nil {
		return 0, err
	}
	return len(s.archivist[id]), nil
}
