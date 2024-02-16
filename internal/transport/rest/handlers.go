package rest

import (
	"flag"
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func StartHandler() {

	const (
		COUNT       = "/"
		CREATE_RULE = "/aggregation_rule/create"
		GET_RULE    = "/aggregation_rules/get"
		PORT        = ":8080"
	)

	addr := flag.String("addr", PORT, "HTTP address")
	flag.Parse()

	s := services.NewStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post(CREATE_RULE, s.CreateAggregationRule)
	r.Get(GET_RULE, s.GetAggregationData)
	r.Post(COUNT, s.CalculateTheAggregated)

	log.Println("Star server on port " + *addr)
	log.Fatal(http.ListenAndServe(*addr, r))
}
