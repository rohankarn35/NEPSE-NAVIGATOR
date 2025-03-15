package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rohankarn35/nepsemarketbot/models"
)

func ParseNepaliDate(date string) (string, error) {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid date format: %s", date)
	}

	_, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid year: %w", err)
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil || month < 1 || month > 12 {
		return "", fmt.Errorf("invalid month: %w", err)
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil || day < 1 || day > 31 { // Note: Adjust max day based on month and year
		return "", fmt.Errorf("invalid day: %w", err)
	}

	// Adjust month index to match NepaliMonths slice (0-based)
	nepaliMonth := models.NepaliMonths[month-1]

	return fmt.Sprintf("%s %d", nepaliMonth.Name, day), nil
}
func ParseEnglishMonth(date string) (string, error) {
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid date format: %s", date)
	}

	_, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", fmt.Errorf("invalid year: %w", err)
	}

	month, err := strconv.Atoi(parts[1])
	if err != nil || month < 1 || month > 12 {
		return "", fmt.Errorf("invalid month: %w", err)
	}

	day, err := strconv.Atoi(parts[2])
	if err != nil || day < 1 || day > 31 { // Note: Adjust max day based on month and year
		return "", fmt.Errorf("invalid day: %w", err)
	}

	// Adjust month index to match NepaliMonths slice (0-based)
	englishMonth := models.EnglishMonths[month-1]

	return fmt.Sprintf("%s %d", englishMonth.Name, day), nil
}
