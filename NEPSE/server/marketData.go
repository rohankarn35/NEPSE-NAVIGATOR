package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"nepseserver/constants"
	"nepseserver/models"
	"net/http"
)

func MarketData() ([]*models.StockData, error) {
	url := constants.STOCK_LIVE_URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch Market data")
	}

	var market models.Market
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&market)
	if err != nil {
		return nil, fmt.Errorf("error decoding data")
	}
	result := make([]*models.StockData, len(market.Result.Stock))
	for i := range market.Result.Stock {
		result[i] = &market.Result.Stock[i]
	}
	log.Print("Market Data Fetched")
	return result, nil

}
