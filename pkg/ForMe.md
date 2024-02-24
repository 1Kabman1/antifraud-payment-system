 

``_**`{
"name":"Amount per client",
"aggregateBy": ["clientId", "bank_card"],
"aggregatedValue": "amount"
}**_
``
````curl -XPOST 'http://127.0.0.1:8080' -H 'Content-Type: appliction/json' -d
'{
"payment_id": 123123,
"client_id": "abasdi-1923",
"payment_method_type": "bank_card",
"payment_method_id": "aopiasdxscnlojcxzoqwe",
"amount": 1000,
"currency": "RUB"
}'`````

===============================================
payment::

`curl -XPOST 'http://127.0.0.1:8080' -H 'Content-Type: appliction/json' -d 
'{
"payment_id": 123123,
"client_id": "abasdi-1923",
"payment_method_type": "bank_card",
"payment_method_id": "aopiasdxscnlojcxzoqwe",
"amount": 1000,
"currency": "RUB"`
}' 



	
			// Добавить разделитель в значение перед кал. хеша (учесть в тесте )+
			// Счетчики должны быть привязаны к правилам нельзя обновить счетчик стороннего правила+
			// убрать многопоточность +
              // Написать тесты учесть колизии совпадения и так далее , так же их прописать в тесте 