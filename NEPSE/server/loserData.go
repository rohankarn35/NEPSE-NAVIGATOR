package server

import (
	"encoding/json"
	"fmt"
	"nepseserver/constants"
	applog "nepseserver/log"
	"nepseserver/models"
	"net/http"
)

func GetMarketMovers(moverType string) ([]*models.MarketMoversData, error) {
	var url string
	if moverType == "gainers" {
		url = constants.TOP_MARKET_MOVERS_URL
	} else if moverType == "losers" {
		url = constants.LOW_MARKET_MOVER_URL
	} else {
		applog.Log(applog.ERROR, "Invalid mover type provided: %s", moverType)
		return nil, fmt.Errorf("invalid url provided")
	}

	applog.Log(applog.INFO, "Fetching market movers from URL: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		applog.Log(applog.ERROR, "Failed to make GET request: %v", err)
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		applog.Log(applog.ERROR, "Failed to fetch data: status code %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}

	var moversData models.MarketMovers
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&moversData); err != nil {
		applog.Log(applog.ERROR, "Failed to decode response body: %v", err)
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	result := make([]*models.MarketMoversData, len(moversData.Result))
	for i := range moversData.Result {
		result[i] = &moversData.Result[i]
	}
	applog.Log(applog.INFO, "Successfully fetched and decoded market movers data")
	return result, nil
}
