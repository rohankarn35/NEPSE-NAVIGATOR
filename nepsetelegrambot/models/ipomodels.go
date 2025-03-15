package models

type IPO struct {
	UniqueSymbol           string `json:"unique_symbol"`
	CompanyName            string `json:"company_name"`
	StockSymbol            string `json:"stock_symbol"`
	ShareRegistrar         string `json:"share_registrar"`
	SectorName             string `json:"sector_name"`
	ShareType              string `json:"share_type"`
	PricePerUnit           string `json:"price_per_unit"`
	Rating                 string `json:"rating"`
	Units                  string `json:"units"`
	MinUnits               string `json:"min_units"`
	MaxUnits               string `json:"max_units"`
	TotalAmount            string `json:"total_amount"`
	OpeningDateAD          string `json:"opening_date_ad"`
	OpeningDateBS          string `json:"opening_date_bs"`
	ClosingDateAD          string `json:"closing_date_ad"`
	ClosingDateBS          string `json:"closing_date_bs"`
	ClosingDateClosingTime string `json:"closing_date_closing_time"`
	Status                 string `json:"status"`
}

type FPO struct {
	UniqueSymbol           string `json:"unique_symbol"`
	CompanyName            string `json:"company_name"`
	StockSymbol            string `json:"stock_symbol"`
	ShareRegistrar         string `json:"share_registrar"`
	SectorName             string `json:"sector_name"`
	ShareType              string `json:"share_type"`
	PricePerUnit           string `json:"price_per_unit"`
	Rating                 string `json:"rating"`
	Units                  string `json:"units"`
	MinUnits               string `json:"min_units"`
	MaxUnits               string `json:"max_units"`
	TotalAmount            string `json:"total_amount"`
	OpeningDateAD          string `json:"opening_date_ad"`
	OpeningDateBS          string `json:"opening_date_bs"`
	ClosingDateAD          string `json:"closing_date_ad"`
	ClosingDateBS          string `json:"closing_date_bs"`
	ClosingDateClosingTime string `json:"closing_date_closing_time"`
	Status                 string `json:"status"`
}

type GetIPOAndFpoAlerts struct {
	IPO []IPO `json:"ipo"`
	FPO []FPO `json:"fpo"`
}

type ResponseData struct {
	GetIPOAndFpoAlerts GetIPOAndFpoAlerts `json:"getIPOAndFpoAlerts"`
}
