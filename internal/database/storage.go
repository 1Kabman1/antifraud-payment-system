package database

type Storage struct {
	idRules    int
	idCounters int
	rules      map[string]interface{}
	counters   map[[16]byte]interface{}
}

func NewStorage() Storage {
	return Storage{
		rules:    make(map[string]interface{}),
		counters: make(map[[16]byte]interface{}),
	}
}
func (s *Storage) SetRule(key string, rule interface{}) {
	s.rules[key] = rule
}

func (s *Storage) GetRulesLen() int {
	return len(s.rules)
}

func (s *Storage) GetRules() map[string]interface{} {

	return s.rules
}

func (s *Storage) GetRule(key string) interface{} {
	return s.rules[key]
}

func (s *Storage) IsRule(key string) bool {
	_, ok := s.rules[key]
	return ok
}

func (s *Storage) GetId() int {
	s.idRules++
	return s.idRules
}

func (s *Storage) IsCounter(key [16]byte) bool {
	_, ok := s.counters[key]
	return ok
}

func (s *Storage) SetIdCounter(key [16]byte, v interface{}) {
	s.counters[key] = v

}

func (s *Storage) GetCounter(key [16]byte) interface{} {
	return s.counters[key]
}

func (s *Storage) GetIdCounter() int {
	s.idCounters++

	return s.idCounters
}
