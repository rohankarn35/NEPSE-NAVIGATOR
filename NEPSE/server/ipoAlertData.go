package server

import (
	"encoding/json"
	"fmt"
	"nepseserver/constants"
	applog "nepseserver/log"
	"nepseserver/models"
	"net/http"
)

func GetIPOAlert(ipoType string) ([]*models.IPO, error) {

	var url string
	if ipoType == "IPO" {
		url = constants.IPO_URL
	} else if ipoType == "FPO" {
		url = constants.FPO_URL
	} else {
		applog.Log(applog.ERROR, "invalid url provided")
		return nil, fmt.Errorf("invalid url provided")
	}

	resp, err := http.Get(url)
	if err != nil {
		applog.Log(applog.ERROR, "failed to make request: %v", err)
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		applog.Log(applog.ERROR, "unexpected status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var ipoAlert models.IPOAlert

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&ipoAlert); err != nil {
		applog.Log(applog.ERROR, "failed to decode response body: %v", err)
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	result := make([]*models.IPO, len(ipoAlert.Result.Data))

	for i := range ipoAlert.Result.Data {
		if ipoAlert.Result.Data[i].Status == "Nearing" {
			ipoAlert.Result.Data[i].Status = "Upcoming"
		}
		result[i] = &ipoAlert.Result.Data[i]
	}
	applog.Log(applog.INFO, "Successfully fetched IPO alerts")
	return result, nil
}
