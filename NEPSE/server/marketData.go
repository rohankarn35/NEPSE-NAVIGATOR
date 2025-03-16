package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"nepseserver/constants"
	applog "nepseserver/log"
	"nepseserver/models"
	"net/http"
)

func MarketData() ([]*models.StockData, error) {
	url := constants.STOCK_LIVE_URL
	resp, err := http.Get(url)
	if err != nil {
		applog.Log(applog.ERROR, "Error fetching data: %v", err)
		return nil, fmt.Errorf("error fetching data")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		applog.Log(applog.ERROR, "Failed to fetch Market data: %v", resp.Status)
		return nil, errors.New("failed to fetch Market data")
	}

	var market models.Market
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&market)
	if err != nil {
		applog.Log(applog.ERROR, "Error decoding data: %v", err)
		return nil, fmt.Errorf("error decoding data")
	}
	result := make([]*models.StockData, len(market.Result.Stock))
	for i := range market.Result.Stock {
		result[i] = &market.Result.Stock[i]
	}
	applog.Log(applog.INFO, "Market Data Fetched")
	return result, nil
}
