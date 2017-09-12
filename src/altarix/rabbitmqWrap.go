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

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func RM_Receive() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func RM_Send() {
	/*Тут иногда надо обращаться по адресу 0.0.0.0:5672/*/
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "hello"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}