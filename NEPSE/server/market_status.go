package server

import (
	"encoding/json"
	"fmt"
	"nepseserver/constants"
	applog "nepseserver/log"
	"nepseserver/models"
	"net/http"
)

func GetMarketStatus() (*models.MarketStatus, error) {
	url := constants.MARKET_STATUS_URL

	resp, err := http.Get(url)
	if err != nil {
		applog.Log(applog.ERROR, "Failed to make request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		applog.Log(applog.ERROR, "Unexpected status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	var marketStatus models.MarketStatus

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&marketStatus); err != nil {
		applog.Log(applog.ERROR, "Failed to decode response body: %v", err)
		return nil, err
	}

	applog.Log(applog.INFO, "Successfully fetched market status")
	return &marketStatus, nil
}
