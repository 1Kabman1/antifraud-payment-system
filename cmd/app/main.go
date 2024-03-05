package main

import (
	"github.com/1Kabman1/antifraud-payment-system/internal/app"
	"log"
)

func main() {

	err := app.StartAntifaud()
	if err != nil {
		log.Fatalln(err)
	}

}
