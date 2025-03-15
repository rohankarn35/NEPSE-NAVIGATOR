package models

type CronJobModel struct {
	UniqueSymbol           string
	StockSymbol            string
	OpeningDateAD          string
	OpeningDateBS          string
	ClosingDateAD          string
	ClosingDateBS          string
	ClosingDateClosingTime string
	Status                 string
}
