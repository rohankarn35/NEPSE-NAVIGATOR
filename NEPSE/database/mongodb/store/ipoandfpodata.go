package store

import (
	"context"
	"fmt"
	"log"
	dbmodels "nepseserver/database/models"
	"nepseserver/server"
	"nepseserver/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func StoreIpoandFpoData(collection *mongo.Collection) {

	ipoData, err := server.GetIPOAlert("IPO")
	if err != nil {
		log.Printf("Error fetching IPO data: %v", err)
		return
	}

	var ipoAlert []dbmodels.IPOAlert
	for _, ipo := range ipoData {
		ipoAlert = append(ipoAlert, dbmodels.IPOAlert{
			UniqueSymbol:           utils.GenerateUniqueSymbol(ipo.ShareType, ipo.StockSymbol),
			CompanyName:            ipo.CompanyName,
			StockSymbol:            ipo.StockSymbol,
			ShareRegistrar:         ipo.ShareRegistrar,
			SectorName:             ipo.SectorName,
			ShareType:              ipo.ShareType,
			PricePerUnit:           ipo.PricePerUnit,
			Rating:                 ipo.Rating,
			Units:                  ipo.Units,
			MinUnits:               ipo.MinUnits,
			MaxUnits:               ipo.MaxUnits,
			TotalAmount:            ipo.TotalAmount,
			OpeningDateAD:          ipo.OpeningDateAD,
			OpeningDateBS:          ipo.OpeningDateBS,
			ClosingDateAD:          ipo.ClosingDateAD,
			ClosingDateBS:          ipo.ClosingDateBS,
			ClosingDateClosingTime: ipo.ClosingDateClosingTime,
			Status:                 ipo.Status,
		})

	}

	fpoData, err := server.GetIPOAlert("FPO")

	if err != nil {
		log.Printf("Error fetching FPO data: %v", err)
		return
	}

	var fpoAlert []dbmodels.IPOAlert
	for _, fpo := range fpoData {
		fpoAlert = append(fpoAlert, dbmodels.IPOAlert{
			UniqueSymbol:           utils.GenerateUniqueSymbol(fpo.ShareType, fpo.StockSymbol),
			CompanyName:            fpo.CompanyName,
			StockSymbol:            fpo.StockSymbol,
			ShareRegistrar:         fpo.ShareRegistrar,
			SectorName:             fpo.SectorName,
			ShareType:              fpo.ShareType,
			PricePerUnit:           fpo.PricePerUnit,
			Rating:                 fpo.Rating,
			Units:                  fpo.Units,
			MinUnits:               fpo.MinUnits,
			MaxUnits:               fpo.MaxUnits,
			TotalAmount:            fpo.TotalAmount,
			OpeningDateAD:          fpo.OpeningDateAD,
			OpeningDateBS:          fpo.OpeningDateBS,
			ClosingDateAD:          fpo.ClosingDateAD,
			ClosingDateBS:          fpo.ClosingDateBS,
			ClosingDateClosingTime: fpo.ClosingDateClosingTime,
			Status:                 fpo.Status,
		})
	}

	ipo := dbmodels.IPO{
		IPO: ipoAlert,
		FPO: fpoAlert,
	}
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		log.Fatalf("Failed to count documents: %v", err)
	}

	if count == 0 {
		_, err := collection.InsertOne(context.TODO(), ipo)
		if err != nil {
			log.Printf("Failed to insert data: %v", err)
		}
	} else {
		_, err := collection.ReplaceOne(context.TODO(), bson.M{}, ipo)
		if err != nil {
			log.Printf("Failed to update data: %v", err)
		}
	}

	if err != nil {
		log.Printf("Error replacing document: %v", err)
	}

	fmt.Print("ipo-fpo updated")

}
