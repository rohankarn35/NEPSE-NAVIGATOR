package cmd

import (
	"log"

	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Print("falied initizaling database")
		return nil
	}

	// Migrate tables in correct order
	if err := db.AutoMigrate(&models.NepseData{}); err != nil {
		log.Printf("failed to migrate NepseData: %v", err)
		return nil
	}
	if err := db.AutoMigrate(&models.CronJob{}); err != nil {
		log.Printf("failed to migrate CronJob: %v", err)
		return nil
	}

	log.Print("Connected to Postgres")
	return db
}
