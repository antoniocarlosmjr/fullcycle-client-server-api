package service

import (
	"context"
	"encoding/json"
	"github.com/fullcycle-client-server-api/server/repository"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/fullcycle-client-server-api/server/models"
)

const TimeoutMaxAwesomeAPI = 200 * time.Millisecond

func GetQuotation() (*models.AwesomeAPIResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutMaxAwesomeAPI)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var awesomeAPIResponse models.AwesomeAPIResponse
	err = json.Unmarshal(body, &awesomeAPIResponse)
	if err != nil {
		return nil, err
	}

	err = repository.SaveQuotation(awesomeAPIResponse)
	if err != nil {
		log.Printf("Error to save quotation: %v", err)
	}

	return &awesomeAPIResponse, nil
}
