package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	CREATE     = "http://127.0.0.1:8080/aggregation_rule/create"
	GET        = "http://127.0.0.1:8080/aggregation_rules/get"
	COMPARABLE = "{\"id\":1,\"Name\":\"Rule1\",\"AggregateBy\":[\"client_id\",\"payment_method_type\",\"payment_method_id\",\"currency\"],\"AggregateValue\":\"count\"}\n{\"id\":2,\"Name\":\"Reule2\",\"AggregateBy\":[\"client_id\",\"payment_method_type\",\"payment_method_id\"],\"AggregateValue\":\"amount\"}\n{\"id\":3,\"Name\":\"Reule3\",\"AggregateBy\":[\"client_id\",\"payment_method_type\",\"payment_method_id\",\"payment_id\"],\"AggregateValue\":\"amount\"}\n"
)

type rule struct {
	Name           string   `json:"Name"`
	AggregateBy    []string `json:"AggregateBy"`
	Amount         int      `json:"AggregatedValue"`
	AggregateValue string   `json:"AggregateValue"`
}

func main() {
	TestPostAndGetForRule()
}

func TestPostAndGetForRule() {
	client := http.Client{} // Создаем клиента

	rules := []rule{
		{
			Name:           "Rule1",
			AggregateBy:    []string{"client_id", "payment_method_type", "payment_method_id", "currency"},
			Amount:         1000,
			AggregateValue: "count",
		},
		{
			Name:           "Reule2",
			AggregateBy:    []string{"client_id", "payment_method_type", "payment_method_id"},
			Amount:         2000,
			AggregateValue: "amount",
		},
		{

			Name:           "Reule3",
			AggregateBy:    []string{"client_id", "payment_method_type", "payment_method_id", "payment_id"},
			Amount:         4000,
			AggregateValue: "amount",
		},
	}

	for _, r := range rules { //   постим на сервер

		jSON, _ := json.Marshal(&r)
		req := bytes.NewReader(jSON)

		client.Post(CREATE, "appliction/json", req)

	}

	resp, err := http.Get(GET) // Получаем ответ от сервера
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Before sorting ", string(body))

	if COMPARABLE != string(body) { // Сравниваем результат
		log.Fatalln("GET requests is not correct")
	}
}
