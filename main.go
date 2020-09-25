package main

import (
	controller "go-mvc-rabbitmq/controller"
	"log"
	"net/http"
)

func main() {
	serveMux := controller.Initialize() //expose to router
	log.Fatal(http.ListenAndServe(":3000", serveMux))
}
