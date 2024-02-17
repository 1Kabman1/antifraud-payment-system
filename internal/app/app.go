package app

import (
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/transport/rest"
)

func StartAntifaud() error {

	return rest.StartHandler()

}
