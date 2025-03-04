package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		fmt.Println("Failed to connect to RabbitMQ", err)
		panic(err)

	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s\n", d.Body)
		}
	}()

	fmt.Println("succesfully connected")
	fmt.Println("waiting for message")
	<-forever

}
