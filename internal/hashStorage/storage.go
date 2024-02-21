package hashStorage

import "errors"

type Storage struct {
	idRules   int
	idCounter int
	rules     map[string]interface{}
	counter   map[[16]byte]interface{}
}

// NewStorage - create a Storage
func NewStorage() Storage {
	return Storage{
		rules:   make(map[string]interface{}),
		counter: make(map[[16]byte]interface{}),
	}
}

// SetRule - adds rule in map
func (s *Storage) SetRule(key string, rule interface{}) {
	s.rules[key] = rule
}

// RuleLen - returns the length of the map
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

func (s *Storage) IsRule(key string) bool {
	_, ok := s.rules[key]
	return ok
}

// IdRules - adds by 1 id and returns
func (s *Storage) IdRules() int {
	s.idRules++
	return s.idRules
}

// IsCounter - Checks if there is a value in the map
func (s *Storage) IsCounter(key [16]byte) bool {
	_, ok := s.counter[key]
	return ok
}

// SetIdCounter - sets id for c
func (s *Storage) SetIdCounter(key [16]byte, v interface{}) {
	s.counter[key] = v

}

// Counter - return Counter
func (s *Storage) Counter(key [16]byte) (error, interface{}) {
	err := errors.New("Key is not correct")

	_, ok := s.counter[key]
	if ok {
		return nil, s.counter[key]
	}
	return err, nil
}

// IdCounter - return id metre
func (s *Storage) IdCounter() int {
	s.idCounter++
	return s.idCounter
}
