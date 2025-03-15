package services

import (
	"strconv"
	"strings"

	"github.com/rohankarn35/nepsemarketbot/models"
)

func ConvertDate(adDate, bsDate string) string {
	// Parse AD date into time.Time
	parsedADDate, _ := ParseEnglishMonth(adDate)

	parsedBSDate, _ := ParseNepaliDate(bsDate)

	// Format AD date to "JAN 20"

	// Combine with BS date
	return parsedBSDate + " ( " + parsedADDate + " )"
}

func BSDateConvert(bsDate string) string {

	// Split the BS date into components
	parts := strings.Split(bsDate, "-")
	year := parts[0]
	monthIndex, _ := strconv.Atoi(parts[1])
	day := parts[2]

	// Get the Nepali month name
	nepaliMonth := models.NepaliMonths[monthIndex-1].Name

	// Format the BS date to "15 Magh 2081"
	formattedBSDate := day + " " + nepaliMonth + " " + year

	return formattedBSDate
}
