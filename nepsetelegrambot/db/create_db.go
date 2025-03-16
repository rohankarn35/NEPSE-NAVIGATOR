package ipodb

import (
	"fmt"

	"github.com/rohankarn35/nepsemarketbot/applog"
	"github.com/rohankarn35/nepsemarketbot/db/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateOrUpdateDB(db *gorm.DB, data models.NepseData) error {
	err := db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "unique_symbol"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"company_name", "stock_symbol", "share_registrar", "sector_name", "share_type",
			"price_per_unit", "rating", "units", "min_units", "max_units", "total_amount",
			"opening_date_ad", "opening_date_bs", "closing_date_ad", "closing_date_bs",
			"closing_date_closing_time", "status", "type", "updated_at",
		}),
	}).Create(&data).Error
	if err != nil {
		applog.Log(applog.ERROR, "Failed to create/update nepse data: %v", err)
		return fmt.Errorf("failed to create/update nepse data: %v", err)
	}
	applog.Log(applog.INFO, "Successfully created/updated nepse data for %s", data.UniqueSymbol)
	return nil
}
