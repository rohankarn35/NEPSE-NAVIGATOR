package server

import (
	"encoding/json"
	"fmt"
	"nepseserver/constants"
	applog "nepseserver/log"
	"nepseserver/models"
	"net/http"
)

// GetIndices fetches the indices data from the API and returns a slice of Index models
func GetIndices() ([]*models.Index, error) {
	url := constants.INDICES_URL
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

	var response models.Response
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		applog.Log(applog.ERROR, "failed to decode response body: %v", err)
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	result := make([]*models.Index, len(response.Result))
	for i, index := range response.Result {
		result[i] = &index
	}

	applog.Log(applog.INFO, "successfully fetched indices data")
	return result, nil
}
