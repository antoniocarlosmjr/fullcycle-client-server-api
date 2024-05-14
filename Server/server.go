package main

import (
	"net/http"

	"github.com/fullcycle-client-server-api/server/controller"
)

func main() {
	http.HandleFunc("/cotacao", controller.GetQuotation)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
