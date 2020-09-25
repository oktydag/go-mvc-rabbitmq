package main

import (
	"log"
	"net/http"

	"github.com/oktydag/go-mvc-rabbitmq/controller"
)

func main() {
	serveMux := controller.Initialize() //expose to router
	log.Fatal(http.ListenAndServe(":3000", serveMux))
}
