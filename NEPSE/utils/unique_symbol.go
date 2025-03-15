package utils

import "strings"

func GenerateUniqueSymbol(shareType string, stockSymbol string) string {
	// Convert shareType to lowercase for case-insensitive comparison
	shareType = strings.ToLower(shareType)

	var suffix string
	switch shareType {
	case "ordinary":
		suffix = "ORD"
	case "migrant workers":
		suffix = "MW"
	case "local":
		suffix = "LO"
	default:
		suffix = "OT"
	}

	return stockSymbol + "_" + suffix
}
