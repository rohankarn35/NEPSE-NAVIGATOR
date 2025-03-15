package server

import (
	"encoding/json"
	"fmt"
	"nepseserver/constants"
	"nepseserver/models"
	"net/http"
)

// GetIndices fetches the indices data from the API and returns a slice of Index models
func GetIndices() ([]*models.Index, error) {
	url := constants.INDICES_URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response models.Response
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	result := make([]*models.Index, len(response.Result))
	for i, index := range response.Result {

		result[i] = &index
	}

	return result, nil
}
