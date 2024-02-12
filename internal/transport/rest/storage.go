package rest

import (
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"sync"
)

type Storage struct {
	IdRules     int
	IdClient    int
	ruleBook    *services.RuleBook
	listClients *services.ListOfClients
	mx          sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		ruleBook:    services.NewRuleBook(),
		listClients: services.NewListOfClients(),
	}
}

func (s *Storage) GiveIdForRules() int {
	s.IdRules++
	return s.IdRules
}

//package rest
//
//import (
//	"github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
//	"sync"
//)
//
//type Storage struct {
//	idStatic int
//	mp       map[[16]byte]*services.Rule
//	mx       sync.Mutex
//}
//
//func NewStorage() *Storage {
//	return &Storage{
//		mp: make(map[[16]byte]*services.Rule),
//	}
//}
