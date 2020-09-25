package controller

import (
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

func receiveMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			failOnError(err, "Failed to connect to RabbitMQ")
			defer conn.Close()

			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			queueName := "mvcpattern"

			msgs, err := ch.Consume(
				queueName,
				"",
				true,
				false,
				false,
				false,
				nil,
			)
			failOnError(err, "Failed to register a consumer")

			messagesFromProvider := ""
			go func() {
				for d := range msgs {
					log.Printf("Received a message: %s", d.Body)
					messagesFromProvider += string(d.Body)
				}
			}()
		}
	}
}
