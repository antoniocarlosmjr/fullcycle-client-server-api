package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fullcycle-client-server-api/client/models"
)

const (
	serverUrl      = "http://localhost:8080/cotacao"
	timeoutRequest = 3000 * time.Millisecond
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutRequest)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, serverUrl, nil)
	if err != nil {
		log.Panicf("Request error: %s", err)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Panicf("Response error: %s", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Panicf("Error reading response body: %s", err)
	}

	var quotation models.USDBRL
	if err = json.Unmarshal(body, &quotation); err != nil {
		log.Panicf("Error unmarshalling response body: %s", err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Panicf("Error creating file: %s", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicf("Error closing file: %s", err)
		}
	}(file)
	data := []byte(quotation.Bid)
	if _, err = file.Write(data); err != nil {
		log.Panicf("Error writing to file: %s", err)
	}
}
