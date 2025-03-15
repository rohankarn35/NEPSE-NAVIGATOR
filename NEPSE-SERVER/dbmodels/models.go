package dbmodels

type MarketMover struct {
	StockCode         string  `bson:"stock_code"`
	Company           string  `bson:"company"`
	TransactionsCount int32   `bson:"transactions_count"`
	HighestPrice      float64 `bson:"highest_price"`
	LowestPrice       float64 `bson:"lowest_price"`
	OpeningPrice      float64 `bson:"opening_price"`
	ClosingPrice      float64 `bson:"closing_price"`
	Turnover          float64 `bson:"turnover"`
	PreviousClose     float64 `bson:"previous_close"`
	PriceChange       float64 `bson:"price_change"`
	PercentageChange  float64 `bson:"percentage_change"`
	TradedVolume      int32   `bson:"traded_volume"`
	TradeDate         string  `bson:"trade_date"`
}

type IPOAndFpoAlert struct {
	Ipo []IPOAlert `bson:"ipo"`
	Fpo []IPOAlert `bson:"fpo"`
}

type MarketMovers struct {
	Gainers []MarketMover `bson:"gainers"`
	Loser   []MarketMover `bson:"losers"`
}

type NepseIndex struct {
	MarketIndex          string  `bson:"market_index"`
	CurrentValue         float64 `bson:"current_value"`
	PreviousClose        float64 `bson:"previous_close"`
	OpeningValue         float64 `bson:"opening_value"`
	PercentageChange     float64 `bson:"percentage_change"`
	PointChange          float64 `bson:"point_change"`
	TotalTurnover        float64 `bson:"total_turnover"`
	TradedVolume         int32   `bson:"traded_volume"`
	MarketCapitalization float64 `bson:"market_capitalization"`
	DailyHigh            float64 `bson:"daily_high"`
	DailyLow             float64 `bson:"daily_low"`
	YearlyHigh           float64 `bson:"yearly_high"`
	YearlyLow            float64 `bson:"yearly_low"`
	Date                 string  `bson:"date"`
}

type IPOAlert struct {
	UniqueSymbol           string `bson:"unique_symbol"`
	CompanyName            string `bson:"company_name"`
	StockSymbol            string `bson:"stock_symbol"`
	ShareRegistrar         string `bson:"share_registrar"`
	SectorName             string `bson:"sector_name"`
	ShareType              string `bson:"share_type"`
	PricePerUnit           string `bson:"price_per_unit"`
	Rating                 string `bson:"rating"`
	Units                  string `bson:"units"`
	MinUnits               string `bson:"min_units"`
	MaxUnits               string `bson:"max_units"`
	TotalAmount            string `bson:"total_amount"`
	OpeningDateAD          string `bson:"opening_date_ad"`
	OpeningDateBS          string `bson:"opening_date_bs"`
	ClosingDateAD          string `bson:"closing_date_ad"`
	ClosingDateBS          string `bson:"closing_date_bs"`
	ClosingDateClosingTime string `bson:"closing_date_closing_time"`
	Status                 string `bson:"status"`
}

type Market struct {
	Symbol           string  `bson:"symbol"`
	Company          string  `bson:"company"`
	TradeVolume      int32   `bson:"trade_volume"`
	High             float64 `bson:"high"`
	Low              float64 `bson:"low"`
	Open             float64 `bson:"open"`
	Close            float64 `bson:"close"`
	TotalTradedValue float64 `bson:"total_traded_value"`
	PrevClose        float64 `bson:"prev_close"`
	PriceChange      float64 `bson:"price_change"`
	PercentChange    float64 `bson:"percent_change"`
	ShareVolume      int32   `bson:"share_volume"`
	LastUpdated      string  `bson:"last_updated"`
}

type Indices struct {
	IndexName        string  `bson:"index_name"`
	IndexValue       float64 `bson:"index_value"`
	PreviousValue    float64 `bson:"previous_value"`
	OpeningValue     float64 `bson:"opening_value"`
	PercentChange    float64 `bson:"percent_change"`
	Difference       float64 `bson:"difference"`
	Turnover         float64 `bson:"turnover"`
	Volume           int32   `bson:"volume"`
	TotalCompanies   int32   `bson:"total_companies"`
	TradedCompanies  int32   `bson:"traded_companies"`
	Transactions     int32   `bson:"transactions"`
	ListedShares     int64   `bson:"listed_shares"`
	MarketCap        float64 `bson:"market_cap"`
	DailyHigh        float64 `bson:"daily_high"`
	DailyLow         float64 `bson:"daily_low"`
	YearlyHigh       float64 `bson:"yearly_high"`
	YearlyLow        float64 `bson:"yearly_low"`
	ReportDate       string  `bson:"report_date"`
	ReportDateString string  `bson:"report_date_string"`
	GainingCompanies int32   `bson:"gaining_companies"`
	LosingCompanies  int32   `bson:"losing_companies"`
	Unchanged        int32   `bson:"unchanged"`
}

type MarketStatus struct {
	IsOpen string `bson:"isOpen"`
}
