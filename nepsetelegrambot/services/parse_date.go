package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rohankarn35/nepsemarketbot/applog"
	"github.com/rohankarn35/nepsemarketbot/models"
)

func ParseNepaliDate(date string) (string, error) {
	applog.Log(applog.INFO, "Parsing Nepali date: %s", date)
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		err := fmt.Errorf("invalid date format: %s", date)
		applog.Log(applog.ERROR, err.Error())
		return "", err
	}

	_, err := strconv.Atoi(parts[0])
	if err != nil {
		applog.Log(applog.ERROR, "Invalid year: %v", err)
		return "", fmt.Errorf("invalid year: %w", err)
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil || month < 1 || month > 12 {
		applog.Log(applog.ERROR, "Invalid month: %v", err)
		return "", fmt.Errorf("invalid month: %w", err)
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil || day < 1 || day > 31 { // Note: Adjust max day based on month and year
		applog.Log(applog.ERROR, "Invalid day: %v", err)
		return "", fmt.Errorf("invalid day: %w", err)
	}

	// Adjust month index to match NepaliMonths slice (0-based)
	nepaliMonth := models.NepaliMonths[month-1]

	result := fmt.Sprintf("%s %d", nepaliMonth.Name, day)
	applog.Log(applog.INFO, "Parsed Nepali date result: %s", result)
	return result, nil
}

func ParseEnglishMonth(date string) (string, error) {
	applog.Log(applog.INFO, "Parsing English date: %s", date)
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		err := fmt.Errorf("invalid date format: %s", date)
		applog.Log(applog.ERROR, err.Error())
		return "", err
	}

	_, err := strconv.Atoi(parts[0])
	if err != nil {
		applog.Log(applog.ERROR, "Invalid year: %v", err)
		return "", fmt.Errorf("invalid year: %w", err)
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil || month < 1 || month > 12 {
		applog.Log(applog.ERROR, "Invalid month: %v", err)
		return "", fmt.Errorf("invalid month: %w", err)
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil || day < 1 || day > 31 { // Note: Adjust max day based on month and year
		applog.Log(applog.ERROR, "Invalid day: %v", err)
		return "", fmt.Errorf("invalid day: %w", err)
	}

	// Adjust month index to match EnglishMonths slice (0-based)
	englishMonth := models.EnglishMonths[month-1]

	result := fmt.Sprintf("%s %d", englishMonth.Name, day)
	applog.Log(applog.INFO, "Parsed English date result: %s", result)
	return result, nil
}
