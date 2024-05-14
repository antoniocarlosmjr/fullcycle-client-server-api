package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fullcycle-client-server-api/server/service"
)

func GetQuotation(w http.ResponseWriter, _ *http.Request) {
	response, err := service.GetQuotation()
	if err != nil {
		log.Printf("Error to get quotation: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response.USDBRL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
