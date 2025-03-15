package server

import (
	"encoding/json"
	"fmt"
	"nepseserver/constants"
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
		return nil, fmt.Errorf("invalid url provided")
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}

	var moversData models.MarketMovers
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&moversData); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	result := make([]*models.MarketMoversData, len(moversData.Result))
	for i := range moversData.Result {
		result[i] = &moversData.Result[i]
	}
	return result, nil
}
