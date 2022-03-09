package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://testUser:test1234@10.211.55.5:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		})
	failOnError(err, "Failed to publish message")
	log.Println("Message sent: ", body)

}

func failOnError(err error, msg string)  {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
