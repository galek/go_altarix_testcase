package main

/*
Поборол ту проблему почему не стартовала служба:

0) Когда будет устанавливаться служба(в инсталяторе), посмотри детальный выхлоп. Там может быть непонятная ошибка. Если вылезла(как у меня), то поздравляю - у тебя будет качественный секс
1) Имя компа не должно иметь латиницу/смешанное название(мое Nick-ПК не канает, как пример.)
2) Ставим Erlang OTP 18(7.0) - C:\erl7.0
3) Ставим RabbitMQ Server 3.6.11 - C:\Rabbit
4) Проверяем переменную ERLANG_HOME, она должна быть C:\erl7.0
5) Проверяем переменную RABBITMQ_BASE, она должна быть C:\Rabbit\rabbitmq_server-3.6.11
6) Удивляемся, почему не работает, даже если выполнить пункт 5
7) Копируем содержимое C:\Rabbit\rabbitmq_server-3.6.11 в C:\Rabbit
8) А вот теперь запустится
9) По умолчанию, висим на локалхосте. Но иногда сервер почему-то был недоступен. Достучаться к нему можно через 0.0.0.0:5762/ . Такая мурда, т.к конфиг по умолчанию пустой
инфа получена через исследование примерного конфига.
10) Отправка и получение теперь работает.

*Замечание, иногда требуется переменную RABBITMQ_BASE присвоить до установки RabbitMQ(по крайней мере про это написано на stackoverflow. У меня так же не взлетало)
*/

/*
Алгоритм, как будем работать:
1) Raw object отправить нельзя. Нужно его сериализовать (см примечание)
2) Формируем очередь сообщений из BD (Нужно поработать над сериализацией MessageIn) как JSON
3) Кидаем сообщения (Raw JSON object)
4) MessageIn ~ctor
5) -> MessageOut
6) WriteToBD MessageOut

"Hi Aluen. RabbitMQ treats message bodies as opaque binary data. If you
want to send objects you'll need to serialise them - if you're only
using Python you could pickle the objects. Otherwise you would need to
write out JSON or XML or similar."


*/

import (
	"database/sql"

	"log"

	"github.com/pquerna/ffjson/ffjson"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		if ISDebug {
			log.Fatalf("%s: %s", msg, err)
		} else {
			if errlog != nil {
				errlog.Fatalf("%s: %s", msg, err)
			} else {
				log.Fatalf("%s: %s", msg, err)
			}
		}
	}
}

func RM_Receive(_name string /*, ref []string*/) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		_name, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	// msgs, err := ch.Get(q.Name, true)
	failOnError(err, "Failed to register a consumer")

	// go-routine
	for d := range msgs {
		go processMessage(d, DB)
	}
}

func processMessage(d amqp.Delivery, _DB *sql.DB) {

	if ISDebug || ISShowSendGetReq {
		log.Printf("[%v] Receive %q", d.DeliveryTag, d.Body)
	}

	inM := MessageIn{}
	ffjson.Unmarshal(d.Body, &inM)

	outM := MessageOut{}
	MessageInToMessageToConverter(&inM, &outM, string(d.Body[:]))

	if ISDebug || ISShowSendGetReq {
		log.Printf("[DEBUG ONLY] processMessage. IN OBJECT, JSON: %s", GenerateJSONIn(inM))
	}

	if ISDebug || ISShowSendGetReq {
		log.Printf("[DEBUG ONLY] processMessage. OUT OBJECT, JSON: %s", GenerateJSONOut(outM))
	}

	WriteMessageToBD(&outM, &_DB)

	d.Ack(false)
}

func RM_Send(_name string, _body string) {
	/*Тут иногда надо обращаться по адресу 0.0.0.0:5672/*/
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		_name, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(_body),
		})

	if ISDebug || ISShowSendGetReq {
		log.Printf(" [x] Sent %s", _body)
	}

	failOnError(err, "Failed to publish a message")
}
