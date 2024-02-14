package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
)

const (
	CREATE     = "http://127.0.0.1:8080/aggregation_rule/create"
	GET        = "http://127.0.0.1:8080/aggregation_rules/get"
	COMPARABLE = "\n\n\n\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\",,,,,,,,,,,,,,,,,,001111222223333445::::::::::::AAAAAABBBNNNVVV[[[]]]aaaaaaaaaaaacccdddeeeeeeeeeeeeeeeeeeggggggggggggggggggiiilllmmmnnnooorrrrrrtttttttttuuuuuuyyy{{{}}}"
)

type rule struct {
	Name           string   `json:"Name"`
	AggregateBy    []string `json:"AggregateBy"`
	Amount         int      `json:"AggregatedValue"`
	AggregateValue string   `json:"AggregateValue"`
}

func TestPostAndGetForRule() {
	client := http.Client{} // Создаем клиента

	for i := 0; i < 3; i++ { // Создаем правила в количестве 3 шт и постим на сервер
		r := rule{
			Name:           strconv.Itoa(i),
			AggregateBy:    []string{strconv.Itoa(i + 1), strconv.Itoa(i + 2), strconv.Itoa(i + 3)},
			Amount:         i,
			AggregateValue: "count",
		}
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

	sort.Slice(body, func(i, j int) bool { // Сортируем так как результат приходит рандомно из-за многопоточной реализации метода GET
		return body[i] < body[j]
	})

	if COMPARABLE != string(body) { // Сравниваем результат
		log.Fatalln("GET requests is not correct")
	}
}

func main() {

	TestPostAndGetForRule()
}
