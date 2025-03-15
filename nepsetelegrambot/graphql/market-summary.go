package dbgraphql

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
	"github.com/rohankarn35/nepsemarketbot/models"
)

func MarketSummary(client *graphql.Client) (*models.MarketSummary, error) {
	req := graphql.NewRequest(`
	
	query{
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
	}

	if err := client.Run(context.Background(), req, &response); err != nil {
		return nil, fmt.Errorf("error in getting the file %w", err)
	}
	return &models.MarketSummary{
		NepseIndex:   response.GetNepseIndex,
		MarketMovers: response.GetMarketMovers,
		Indices:      response.GetIndices,
	}, nil

}
