package dbgraphql

import (
	"context"

	"github.com/machinebox/graphql"
	"github.com/rohankarn35/nepsemarketbot/models"
)

func GetIPOFPODetails(client *graphql.Client) ([]models.IPO, []models.FPO, error) {
	req := graphql.NewRequest(`
		query {
			getIPOAndFpoAlerts {
				ipo {
				unique_symbol
					company_name
					stock_symbol
					share_registrar
					sector_name
					share_type
					price_per_unit
					rating
					units
					min_units
					max_units
					total_amount
					opening_date_ad
					opening_date_bs
					closing_date_ad
					closing_date_bs
					closing_date_closing_time
					status
				}
				fpo {
				unique_symbol
					company_name
					stock_symbol
					share_registrar
					sector_name
					share_type
					price_per_unit
					rating
					units
					min_units
					max_units
					total_amount
					opening_date_ad
					opening_date_bs
					closing_date_ad
					closing_date_bs
					closing_date_closing_time
					status
				}
			}
		}
	`)

	var respData models.ResponseData

	if err := client.Run(context.Background(), req, &respData); err != nil {
		return nil, nil, err
	}

	return respData.GetIPOAndFpoAlerts.IPO, respData.GetIPOAndFpoAlerts.FPO, nil
}
