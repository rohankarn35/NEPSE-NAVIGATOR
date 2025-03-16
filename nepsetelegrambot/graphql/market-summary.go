package dbgraphql

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/rohankarn35/nepsemarketbot/applog" // Import the applog package
	"github.com/rohankarn35/nepsemarketbot/models"
)

func MarketSummary(client *graphql.Client) (*models.MarketSummary, error) {
	applog.Log(applog.INFO, "Starting MarketSummary function")

	req := graphql.NewRequest(`
  query{

   getMarketStatus{
  isMarketOpen
}
    getNepseIndex{
      index_value
      percent_change
      difference
      turnover
      volume
      as_of_date
    }
    getMarketMovers(top:5){
      gainers{
        stock_symbol
        difference_rs
        percent_change
      }
      losers{
        stock_symbol
        difference_rs
        percent_change
      }
    }
    getIndices(top:3){
      index_name
      percent_change
      difference
    }
  }
  `)

	var response struct {
		GetNepseIndex   models.NepseIndex   `json:"getNepseIndex"`
		GetMarketMovers models.MarketMovers `json:"getMarketMovers"`
		GetIndices      []models.Indices    `json:"getIndices"`
		GetMarketStatus models.MarketStatus `json:"getMarketStatus"`
	}

	applog.Log(applog.DEBUG, "Sending GraphQL request")
	if err := client.Run(context.Background(), req, &response); err != nil {
		applog.Log(applog.ERROR, "Error in getting the file: %v", err)
		return nil, fmt.Errorf("error in getting the file %w", err)
	}
	fmt.Print("the market status is ", response.GetMarketStatus)

	applog.Log(applog.INFO, "Successfully retrieved market summary data")
	return &models.MarketSummary{
		NepseIndex:   response.GetNepseIndex,
		MarketMovers: response.GetMarketMovers,
		Indices:      response.GetIndices,
		MarketStatus: response.GetMarketStatus,
	}, nil
}
