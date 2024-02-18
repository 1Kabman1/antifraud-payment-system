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
		COUNT       = "/"
		CREATE_RULE = "/aggregation_rule/create"
		GET_RULE    = "/aggregation_rules/get"
		PORT        = ":8080"
	)

	addr := flag.String("addr", PORT, "HTTP address")
	flag.Parse()

	h := services.Handler{}
	h.SetStorage()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post(CREATE_RULE, h.CreateAggregationRule)
	r.Get(GET_RULE, h.GetAggregationData)
	r.Post(COUNT, h.CalculateTheAggregated)

	log.Println("Star server on port " + *addr)
	return http.ListenAndServe(*addr, r)
}
