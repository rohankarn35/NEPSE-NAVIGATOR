package ipodb

import (
	"fmt"
	"log"

	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/gorm"
)

func UpdateStatus(db *gorm.DB, uniqueSymbol string, status string) error {
	result := db.Model(&models.NepseData{}).Where("unique_symbol=?", uniqueSymbol).Update("status", status)

	if result.Error != nil {
		log.Fatalf("Error updating IPO status: %v\n", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		fmt.Printf("No records updated for StockSymbol: %s\n", uniqueSymbol)
	} else {
		fmt.Println("IPO status updated successfully!")
	}
	return nil
}
