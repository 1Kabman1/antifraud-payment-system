package rest

import (
	"flag"
	"github.com/1Kabman1/Antifraud-payment-system.git/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func StartHandler() error {

	const (
		Count      = "/"
		CreateRule = "/aggregation_rule/create"
		GetRule    = "/aggregation_rules/get"
		Port       = ":8080"
	)

	addr := flag.String("addr", Port, "HTTP address")
	flag.Parse()

	h := services.Handlers{}
	h.SetStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post(CreateRule, h.CreateAggregationRule)
	r.Get(GetRule, h.AggregationData)
	r.Post(Count, h.CalculateTheAggregated)

	log.Println("Star server on port " + *addr)
	return http.ListenAndServe(*addr, r)
}
