package ipodb

import (
	"github.com/rohankarn35/nepsemarketbot/applog"
	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/gorm"
)

func CheckAndUpdateIPOStatus(db *gorm.DB, uniqueSymbol string, status string) bool {
	var nepseData models.NepseData
	result := db.Where("unique_symbol = ?", uniqueSymbol).First(&nepseData)

	// If no records found, it's a new IPO
	if result.Error == gorm.ErrRecordNotFound {
		return true
	} else if result.Error != nil {
		applog.Log(applog.ERROR, "Error querying IPO status: %v", result.Error)
		return false
	}

	applog.Log(applog.INFO, "Existing Status: %s, New Status: %s, IPO Name: %s", nepseData.Status, status, uniqueSymbol)

	return nepseData.Status != status
}
