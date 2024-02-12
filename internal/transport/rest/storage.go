package rest

import (
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
)

type Storage struct {
	idStatic int
	rules    map[[16]byte]*services.Rule
}

func NewStorage() *Storage {
	return &Storage{
		rules: make(map[[16]byte]*services.Rule),
	}
}
