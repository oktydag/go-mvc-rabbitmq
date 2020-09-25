package viewmodel

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type RabbitMQProviderRequest struct {
	QueueName string      `json:"name"`
	Body      interface{} `json:"body"`
}

type RabbitMQConsumerResponse struct {
	QueueName string      `json:"name"`
	Body      interface{} `json:"body"`
}
