#  Начало работы:
Для начала работы Вам нужно клонировать репозиторий с сайта GitHub с помощью команды 
**`git clone git@github.com:1Kabman1/antifraud-payment-system.git`**
После чего вам следует произвести тестирование приложения с помощью команды **`go test ./...`** если тестирование
пройдет успешно, запустите приложение с помощью команды **`go run .`** , обязательно убедитесь что Вы находитесь в 
директории приложения.

## POST// создание  Rule
Чтобы создать правило агрегации, нужно отправить POST запрос на url "http://127.0.0.1:8080/aggregation_rule/create"

   **`curl -XPOST -v  'http://127.0.0.1:8080/aggregation_rule/create' -H 'Content-Type: appliction/json' -d {}`**

в формате JSON:

Пример №1:
```json
{
    "Name": "<уникальное имя>",
    "AggregateBy": "<агрегируемые свойства>",
    "AggregateValue": "<сумма или счетчик>"
}
```
Пример №2:

```json
{
    "Name": "Amount per client",
    "AggregateBy": ["clientId", "bank_card"],
    "AggregateValue": "amount"
} 
```
В следствии программа создаст правила агрегации с уникальным id. Имя правила Вы задаете сами, имена могут быть 
одинаковые и агрегируемые свойства, но при этом правила все равно будут уникальны по отношению друг к другу, 
уникальность правилам придает уникальный id для каждого правила. 

Пример:
```json
{
  "1": {
    "AggregationRuleId": 1,
    "Name": "Amount per client",
    "AggregateBy": [
      "clientId",
      "bank_card"
    ],
    "AggregateValue": "amount"
  },
  "2": {
    "AggregationRuleId": 2,
    "Name": "Amount per client",
    "AggregateBy": [
      "clientId",
      "bank_card"
    ],
    "AggregateValue": "amount"
  }
}
```

## GET// запрос Rule 
Чтобы получить данные по созданным правилам, нужно отправить GET запрос на 
url "http://127.0.0.1:8080/aggregation_rules/get".

**`curl -XGET -v 'http://127.0.0.1:8080/aggregation_rules/get'`**

В следствии программа вернет ответ в JSON формате. (! Программа возвращает полный список правил)

Пример:

```json
{
  "1": {
    "AggregationRuleId": 1,
    "Name": "Amount per client",
    "AggregateBy": [
      "clientId",
      "bank_card"
    ],
    "AggregateValue": "amount"
  },
  "2": {
    "AggregationRuleId": 2,
    "Name": "Amount per client",
    "AggregateBy": [
      "clientId",
      "bank_card"
    ],
    "AggregateValue": "amount"
  }
}
```

## POST // Отслеживание агрегируемого  

Чтобы начать отслеживать агрегируемое на основании созданных правил, отправьте POST запрос на 
url "http://127.0.0.1:8080/register"

**`curl -POST -v  'http://127.0.0.1:8080/register' -H 'Content-Type: appliction/json' -d '{}'`**

Пример:

```json
{
"payment_id": 123123,
"client_id": "abasdi-1923",
"payment_method_type": "bank_card",
"payment_method_id": "aopiasdxscnlojcxzoqwe",
"amount": 1000,
"currency": "RUB"
}
```
Программа проверит данное агрегируемое по всем правилам, а так же зафиксирует "count" OR "amount" в специально созданном
уникальном счетчике.
 