package ipodb

import (
	"fmt"

	"github.com/rohankarn35/nepsemarketbot/applog"
	gorm_model "github.com/rohankarn35/nepsemarketbot/db/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func StoreCron(db *gorm.DB, cron gorm_model.CronJob) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Upsert NepseData
	if err := tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "unique_symbol"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"company_name", "stock_symbol", "share_registrar", "sector_name", "share_type",
			"price_per_unit", "rating", "units", "min_units", "max_units", "total_amount",
			"opening_date_ad", "opening_date_bs", "closing_date_ad", "closing_date_bs",
			"closing_date_closing_time", "status", "type", "updated_at",
		}),
	}).Create(&cron.NepseData).Error; err != nil {
		tx.Rollback()
		applog.Log(applog.ERROR, "Failed to upsert nepse data: %v", err)
		return fmt.Errorf("failed to upsert nepse data: %v", err)
	}

	// Check if CronJob already exists; if so, skip or update based on your needs
	var existingCron gorm_model.CronJob
	if err := tx.Where("unique_symbol = ?", cron.UniqueSymbol).First(&existingCron).Error; err == nil {
		// CronJob exists; decide whether to update or skip
		applog.Log(applog.INFO, "Cron job for %s already exists, skipping", cron.UniqueSymbol)
		tx.Commit()
		return nil
	} else if err != gorm.ErrRecordNotFound {
		tx.Rollback()
		applog.Log(applog.ERROR, "Failed to check existing cron job: %v", err)
		return fmt.Errorf("failed to check existing cron job: %v", err)
	}

	// Create CronJob if it doesn't exist
	if err := tx.Create(&cron).Error; err != nil {
		tx.Rollback()
		applog.Log(applog.ERROR, "Failed to store cron job: %v", err)
		return fmt.Errorf("failed to store cron job: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		applog.Log(applog.ERROR, "Failed to commit transaction: %v", err)
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	applog.Log(applog.INFO, "Successfully stored cron job for %s", cron.UniqueSymbol)
	return nil
}

func ReadCron(db *gorm.DB) ([]gorm_model.CronJob, error) {
	var cron []gorm_model.CronJob

	if err := db.Preload("NepseData").Find(&cron).Error; err != nil {
		applog.Log(applog.ERROR, "failed to load data: %v", err)
		return nil, fmt.Errorf("failed to load data")
	}

	applog.Log(applog.INFO, "read all the documents")
	return cron, nil
}
