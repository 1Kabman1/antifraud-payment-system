package services

type Storage struct {
	idRules    int
	idCounters int
	rules      map[string]rule
	counters   map[[16]byte]counter
}

func NewStorage() Storage {
	return Storage{
		rules:    make(map[string]rule),
		counters: make(map[[16]byte]counter),
	}
}
