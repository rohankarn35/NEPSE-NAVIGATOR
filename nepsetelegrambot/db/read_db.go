package ipodb

import (
	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/gorm"
)

// func ReadDB(db *pgxpool.Pool) {

// }
func ReadDB(db *gorm.DB, status, stockType string) []models.NepseData {
	var nepseData []models.NepseData
	if err := db.Where("status = ? AND type = ?", status, stockType).Find(&nepseData).Error; err != nil {
		// handle error
		return nil
	}
	// process nepseData as needed
	return nepseData
}
