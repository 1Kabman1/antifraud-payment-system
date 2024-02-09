package app

import (
	"flag"
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/transport/rest"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func StartAntifaud() {

	const (
		CREATE = "/aggregation_rule/CREATE"
		GET    = "/aggregation_rules/GET"
		START  = "Star server on port "
		PORT   = ":8080"
		AD     = "HTTP address"
		ADDR   = "addr"
	)

	addr := flag.String(ADDR, PORT, AD)
	flag.Parse()

	s := rest.NewStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post(CREATE, s.CreateAggregationRoull)
	r.Get(GET, s.GetAggregationData)

	log.Println(START + *addr)
	log.Fatal(http.ListenAndServe(*addr, r))
}
