package ipodb

import (
	"fmt"
	"log"

	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateDB(db *gorm.DB, nepse models.NepseData) error {
	result := db.Model(&models.NepseData{}).Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "unique_symbol"}},
		DoUpdates: clause.AssignmentColumns(
			[]string{
				"status",
				"opening_date_ad",
				"opening_date_bs",
				"closing_date_ad",
				"closing_date_bs",
			},
		),
		Where: clause.Where{
			Exprs: []clause.Expression{
				clause.Neq{Column: "nepse_data.status", Value: nepse.Status},
			},
		},
	}).Create(&nepse)
	if result.Error != nil {
		log.Fatalf("Error inserting or updating IPO: %v\n", result.Error)
		return result.Error
	}

	fmt.Println("IPO created or updated successfully!")
	return nil
}
