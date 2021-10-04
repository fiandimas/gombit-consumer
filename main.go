package main

import (
	"gombit-consumer/config"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial(config.URL)
	__(err)
	defer conn.Close()

	ch, err := conn.Channel()
	__(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"dxdlasd",
		false,
		false,
		false,
		false,
		nil,
	)
	__(err)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	__(err)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func __(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
