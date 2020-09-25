package controller

import (
	"net/http"
)

func Initialize() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/", giveFive())
	serveMux.HandleFunc("/publish", publishMessage())
	serveMux.HandleFunc("/receive", receiveMessage())

	return serveMux
}
