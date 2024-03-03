package hashStorage

import (
	"errors"
	"strconv"
)

type Storage struct {
	rules     map[string]interface{}
	counter   map[[16]byte]interface{}
	archivist map[int][]int
}

// NewStorage - create a Storage
func NewStorage() Storage {
	return Storage{
		rules:     make(map[string]interface{}),
		counter:   make(map[[16]byte]interface{}),
		archivist: make(map[int][]int, 5),
	}
}

// SetRule - set rule in map
func (s *Storage) SetRule(key string, rule interface{}) {
	s.rules[key] = rule
}

// RulesLen - returns the length of the map
func (s *Storage) RulesLen() int {
	return len(s.rules)
}

// Rules - returns rules
func (s *Storage) Rules() map[string]interface{} {
	return s.rules
}

// Rule - returns a rule
func (s *Storage) Rule(key string) (error, interface{}) {
	err := errors.New("Key is not correct")

	_, ok := s.rules[key]
	if ok {
		return nil, s.rules[key]
	}

	return err, nil
}

// IsRule - return bool
func (s *Storage) IsRule(key string) bool {
	_, ok := s.rules[key]
	return ok
}

// IsCounter - Checks if there is a value in the map
func (s *Storage) IsCounter(key [16]byte) bool {
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
