package server

import (
	"encoding/json"
	"fmt"
	"nepseserver/constants"
	"nepseserver/models"
	"net/http"
)

func GetMarketStatus() (*models.MarketStatus, error) {
	url := constants.MARKET_STATUS_URL

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	var marketStatus models.MarketStatus

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&marketStatus); err != nil {
		return nil, fmt.Errorf("failed to decode response body :%w", err)
	}

	return &marketStatus, nil

}
