package server

import (
	"encoding/json"
	"errors"
	"nepseserver/constants"
	applog "nepseserver/log"
	"nepseserver/models"
	"net/http"
)

func FetchNepseData() (*models.NepseIndex, error) {
	applog.Log(applog.INFO, "Fetching Nepse data from URL: %s", constants.INDEX_LIVE_URL)
	resp, err := http.Get(constants.INDEX_LIVE_URL)
	if err != nil {
		applog.Log(applog.ERROR, "Error fetching Nepse data: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		applog.Log(applog.ERROR, "Failed to fetch Nepse data, status code: %d", resp.StatusCode)
		return nil, errors.New("failed to fetch Nepse data")
	}

	var nepseData models.NepseLive
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&nepseData)
	if err != nil {
		applog.Log(applog.ERROR, "Error decoding Nepse data: %v", err)
		return nil, err
	}

	// Extract only Nepse data
	for _, index := range nepseData.Result {
		if index.IndexName == "Nepse" {
			applog.Log(applog.INFO, "Nepse data found: %+v", index)
			return &index, nil
		}
	}

	applog.Log(applog.WARN, "Nepse data not found")
	return nil, errors.New("nepse data not found")
}
