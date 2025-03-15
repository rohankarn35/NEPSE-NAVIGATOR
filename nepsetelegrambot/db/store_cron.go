package ipodb

import (
	"fmt"
	"log"

	gorm_model "github.com/rohankarn35/nepsemarketbot/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func StoreCron(db *gorm.DB, cron gorm_model.CronJob) error {
	// First, handle the NepseData (upsert it)
	if err := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "unique_symbol"}},
		DoNothing: true, // If conflict, do nothing (don’t update existing record)
	}).Create(&cron.NepseData).Error; err != nil {
		return fmt.Errorf("failed to create nepse data: %v", err)
	}

	// Step 2: Create the CronJob, linking to the existing or new NepseData
	if err := db.Create(&cron).Error; err != nil {
		return fmt.Errorf("failed to store cron job: %v", err)
	}

	return nil
}

func ReadCron(db *gorm.DB) ([]gorm_model.CronJob, error) {
	var cron []gorm_model.CronJob

	if err := db.Preload("NepseData").Find(&cron).Error; err != nil {
		return nil, fmt.Errorf("failed to load data")
	}
	log.Print("read all the documents", cron)
	return cron, nil
}
