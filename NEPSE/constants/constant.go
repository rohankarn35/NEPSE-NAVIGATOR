package constants

import (
	"os"
)

var (
	DB_URL                string
	INDICES_URL           string
	IPO_URL               string
	TOP_MARKET_MOVERS_URL string
	LOW_MARKET_MOVER_URL  string
	STOCK_LIVE_URL        string
	INDEX_LIVE_URL        string
	FPO_URL               string
	MARKET_STATUS_URL     string
)

func InitConstant() {

	DB_URL = os.Getenv("DB_URL")
	INDICES_URL = os.Getenv("INDICES_URL")
	IPO_URL = os.Getenv("IPO_URL")
	FPO_URL = os.Getenv("FPO_URL")
	TOP_MARKET_MOVERS_URL = os.Getenv("TOP_MARKET_MOVERS_URL")
	LOW_MARKET_MOVER_URL = os.Getenv("Low_MARKET_MOVER_URL")
	STOCK_LIVE_URL = os.Getenv("STOCK_LIVE_URL")
	INDEX_LIVE_URL = os.Getenv("INDEX_LIVE_URL")
	MARKET_STATUS_URL = os.Getenv("MARKET_STATUS_URL")
}
