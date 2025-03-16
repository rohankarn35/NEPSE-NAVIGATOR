package ipodb

import (
	"github.com/rohankarn35/nepsemarketbot/applog" // Import the applog package
	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/gorm"
)

func ReadDB(db *gorm.DB, status, stockType string) []models.NepseData {
	var nepseData []models.NepseData
	if err := db.Where("status = ? AND type = ?", status, stockType).Find(&nepseData).Error; err != nil {
		applog.Log(applog.ERROR, "Failed to read from database: %v", err) // Log the error
		return nil
	}
	applog.Log(applog.INFO, "Successfully read from database: %d records found", len(nepseData)) // Log the success
	// process nepseData as needed
	return nepseData
}
