package ipodb

import (
	"fmt"
	"log"

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
		log.Fatalf("Error querying IPO status: %v\n", result.Error)
	}

	log.Printf("Existing Status: %s, New Status: %s, IPO Name: %s\n", nepseData.Status, status, uniqueSymbol)

	fmt.Print(nepseData.Status != status)

	return nepseData.Status != status
}
