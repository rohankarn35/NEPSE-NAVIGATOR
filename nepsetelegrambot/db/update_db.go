package ipodb

import (
	"github.com/rohankarn35/nepsemarketbot/applog"
	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/gorm"
)

func UpdateStatus(db *gorm.DB, uniqueSymbol string, status string) error {
	result := db.Model(&models.NepseData{}).Where("unique_symbol=?", uniqueSymbol).Update("status", status)

	if result.Error != nil {
		applog.Log(applog.ERROR, "Error updating IPO status: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		applog.Log(applog.WARN, "No records updated for StockSymbol: %s", uniqueSymbol)
	} else {
		applog.Log(applog.INFO, "IPO status updated successfully!")
	}
	return nil
}
