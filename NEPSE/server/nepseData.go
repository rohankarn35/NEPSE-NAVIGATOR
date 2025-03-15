package server

import (
	"encoding/json"
	"errors"
	"nepseserver/constants"
	"nepseserver/models"
	"net/http"
)

func FetchNepseData() (*models.NepseIndex, error) {
	resp, err := http.Get(constants.INDEX_LIVE_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch Nepse data")
	}

	var nepseData models.NepseLive
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&nepseData)
	if err != nil {
		return nil, err
	}

	// Extract only Nepse data
	for _, index := range nepseData.Result {
		if index.IndexName == "Nepse" {
			return &index, nil
		}
	}

	return nil, errors.New("nepse data not found")
}
