package models

type NepseIndex struct {
	IndexValue    float64 `json:"index_value"`
	PercentChange float64 `json:"percent_change"`
	Difference    float64 `json:"difference"`
	Turnover      float64 `json:"turnover"`
	Volume        int     `json:"volume"`
	AsOfDate      string  `json:"as_of_date"`
}

type MarketMover struct {
	StockSymbol   string  `json:"stock_symbol"`
	Amount        float64 `json:"difference_rs"`
	PercentChange float64 `json:"percent_change"`
}

type Indices struct {
	IndexName     string  `json:"index_name"`
	PercentChange float64 `json:"percent_change"`
	Difference    float64 `json:"difference"`
}

type MarketSummary struct {
	NepseIndex   NepseIndex   `json:"getNepseIndex"`
	MarketMovers MarketMovers `json:"getMarketMovers"`
	Indices      []Indices    `json:"getIndices"`
	MarketStatus MarketStatus `json:"getMarketStatus"`
}

type MarketMovers struct {
	Gainers []MarketMover `json:"gainers"`
	Losers  []MarketMover `json:"losers"`
}

type MarketStatus struct {
	IsMarketOpen bool `json:"isMarketOpen"`
}
