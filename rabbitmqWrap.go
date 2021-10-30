package main

import (
	"github.com/streadway/amqp"
	"log"
)

func Consumer(Name string) error {
	// TODO: вынести это в конфиг
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		Name, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		// TODO: проверить есть ли паника тут
		log.Fatalf("Failed to declare a queue: %s", err)
		return err
	}

	// TODO: Это же консюмер, посмотри где producer
	messages, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack // Автоматическое подтверждение что получили сообщение
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
		return err
	}

	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			// For example, show received message in a console.
			log.Printf(" > Received message: %s\n", message.Body)
			// TODO: replace it
			//processMessage(message, nil)
		}
	}()

	<-forever

	return nil
}

/*
func processMessage(d amqp.Delivery) error {

	log.Printf("[%v] Receive %q", d.DeliveryTag, d.Body)

	inM := MessageIn{}
	err := json.Unmarshal(d.Body, &inM)
	if err != nil {
		return err
	}

	//outM := MessageOut{}

	// Тут логика
	// Тут подтверждение
	d.Ack(false)

	return nil
}
*/
