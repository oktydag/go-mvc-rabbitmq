package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/oktydag/go-mvc-rabbitmq/viewmodel"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func publishMessage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
			failOnError(err, "Failed to connect to RabbitMQ")
			defer conn.Close()

			ch, err := conn.Channel()
			failOnError(err, "Failed to open a channel")
			defer ch.Close()

			queueName := "mvcpattern"
			q, err := ch.QueueDeclare(
				queueName,
				false,
				false,
				false,
				false,
				nil,
			)
			failOnError(err, "Failed to declare a queue")

			body := "Hey, This is a Provider whose name is mvcpattern"
			err = ch.Publish(
				"",
				q.Name,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(body),
				})
			log.Printf(" [x] Sent %s", body)
			failOnError(err, "Failed to publish a message")

			responseMessage := viewmodel.RabbitMQProviderRequest{
				QueueName: queueName,
				Body:      body,
			}

			json.NewEncoder(w).Encode(responseMessage)
		}
	}
}
