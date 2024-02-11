package rest

import (
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"sync"
)

type Storage struct {
	idStatic int
	mp       map[[16]byte]*services.Rule
	mx       sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		mp: make(map[[16]byte]*services.Rule),
	}
}
