#  Начало работы:
Запускаем функцию `StartAntifraud()`

## POST// создание  Rule

Чтобы создать правило агрегации нужно отправить POST запрос на url "http://127.0.0.1:8080/aggregation_rule/create" в формате JSON:

```json
{
    "Name": "<уникальное имя>",
    "AggregateBy": "<агрегируемые свойства>",
    "AggregateValue": "<сумма или счетчик>"
}
```

Пример:

```json
{
    "Name": "Amount per client",
    "AggregateBy": ["clientId", "bank_card"],
    "AggregateValue": "amount"
}
```

В следствии чего программа создаст правила агрегации с уникальными именами, уникальным id. Если правило уже существует, программа отправит ответ 409.

## GET// запрос Rule 
Чтобы получить данные по созданным правилам, нужно отправить GET запрос на url "http://127.0.0.1:8080/aggregation_rules/get".

Пример:
 
```json
{
    "id":1,
    "Name":"Amount per client",
    "AggregateBy":["clientId", "bank_card"],
    "AggregateValue":"amount"
}
```

В следствии чего программа вернет ответ в JSON формате. (! Программа возвращает полный список правил)


## POST // Отслеживание агрегируемого  

Чтобы начать отслеживать аргументированное на основании созданных правил отправьте POST  запрос на url  "http://127.0.0.1:8080/"

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
Программа проверит данное агрегируемое по всем правилам, а так же запишет.
 