package controller

import (
	"encoding/json"

	"net/http"

	"github.com/oktydag/go-mvc-rabbitmq/viewmodel"
)

func giveFive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := viewmodel.Response{
				Code: http.StatusOK,
				Body: "Hey ! Give me Five",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
