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
	COMPARABLE = "\n\n\n\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\"\",,,,,,,,,,,,,,,,,,,,,000111111122222233333445:::::::::::::::AAAAAABBBCCCNNNVVV[[[]]]aaaaaaaaaaaaddddddeeeeeeeeeeeeeeeeeeggggggggggggggggggiiilllmmmnnnooorrrrrrtttttttttuuuuuuyyy{{{}}}"
)

type rule struct {
	Name        string   `json:"Name"`
	AggregateBy []string `json:"AggregateBy"`
	Amount      int      `json:"AggregatedValue"`
}

func main() {

	client := http.Client{} // Создаем клиента

	for i := 0; i < 3; i++ { // Создаем правила в количестве 3 шт и постим на сервер
		r := rule{
			Name:        strconv.Itoa(i),
			AggregateBy: []string{strconv.Itoa(i + 1), strconv.Itoa(i + 2), strconv.Itoa(i + 3)},
			Amount:      i,
		}
		tempReq, _ := json.Marshal(&r)
		req := bytes.NewReader(tempReq)

		client.Post(CREATE, "appliction/json", req)

	}

	r := rule{ // Создаем еще одно правило которое уже существует "Name":"2","AggregateBy":["3","4","5","2"],"AggregatedValue":2,"Count":1

		Name:        strconv.Itoa(2),
		AggregateBy: []string{strconv.Itoa(3), strconv.Itoa(4), strconv.Itoa(5)},
		Amount:      1,
	}
	tempReq, _ := json.Marshal(&r)
	request := bytes.NewReader(tempReq)

	client.Post("http://127.0.0.1:8080/aggregation_rule/create", "appliction/json", request)

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

	sort.Slice(body, func(i, j int) bool { // Сортируем так как результат приходит рандомно из-за многопоточной реализации метода
		return body[i] < body[j]
	})

	if COMPARABLE != string(body) { // Сравниваем результат
		log.Fatalln("GET requests is not correct")
	}

}
