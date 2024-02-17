package main

import (
	antifraud "github.com/1Kabman1/Antifraud-payment-system.git/internal/app"
	"log"
)

func main() {
	err := antifraud.StartAntifaud()
	if err != nil {
		log.Fatalln(err)
	}

}
