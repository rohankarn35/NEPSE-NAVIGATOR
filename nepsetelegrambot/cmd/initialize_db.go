package cmd

import (
	"github.com/rohankarn35/nepsemarketbot/applog"
	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDb(dbUrl string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		applog.Log(applog.ERROR, "failed initializing database")
		return nil
	}

	// Migrate tables in correct order
	if err := db.AutoMigrate(&models.NepseData{}); err != nil {
		applog.Log(applog.ERROR, "failed to migrate NepseData: %v", err)
		return nil
	}
	if err := db.AutoMigrate(&models.CronJob{}); err != nil {
		applog.Log(applog.ERROR, "failed to migrate CronJob: %v", err)
		return nil
	}

	applog.Log(applog.INFO, "Connected to Postgres")
	return db
}
