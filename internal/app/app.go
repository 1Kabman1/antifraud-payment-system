package app

import (
	"github.com/1Kabman1/antifraud-payment-system/internal/transport/rest"
)

func StartAntifaud() error {

	return rest.StartHandler()

}
