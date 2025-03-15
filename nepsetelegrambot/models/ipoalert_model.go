package models

type IPOAlertModel struct {
	UniqueSymbol           string
	CompanyName            string
	StockSymbol            string
	ShareRegistrar         string
	SectorName             string
	ShareType              string
	PricePerUnit           string
	Rating                 string
	Units                  string
	MinUnits               string
	MaxUnits               string
	TotalAmount            string
	OpeningDateAD          string
	OpeningDateBS          string
	ClosingDateAD          string
	ClosingDateBS          string
	ClosingDateClosingTime string
	Status                 string
	Type                   string
}
