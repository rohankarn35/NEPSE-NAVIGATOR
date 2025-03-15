package cronjobs

import (
	"fmt"
	"log"
	"nepseserver/database/mongodb/store"
	"nepseserver/server"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func (cr *CronJob) ScheduleDailyMarketJobs(mongodatabase *mongo.Database) {
	store.StoreIndicesData(mongodatabase.Collection("indices-data"))
	store.StoreNepseData(mongodatabase.Collection("nepse-data"))
	store.MarketMovers(mongodatabase.Collection("marketmovers"))

	_, err := cr.c.AddFunc("50 14 * * 0-4", func() {
		fmt.Println("Checking market status at 2:50 PM.")
		marketStatus, _ := server.GetMarketStatus()
		if !strings.EqualFold(marketStatus.IsOpen, "OPEN") {
			fmt.Println("Market is closed, skipping the job for today.")
			return
		}
		_, err := cr.c.AddFunc("8 15 * * 0-4", func() {
			fmt.Println("Cron job with ID 1 is scheduled.")
			// Run all three functions
			store.StoreIndicesData(mongodatabase.Collection("indices-data"))
			store.StoreNepseData(mongodatabase.Collection("nepse-data"))
			store.MarketMovers(mongodatabase.Collection("marketmovers"))
		})
		if err != nil {
			fmt.Printf("Error scheduling market jobs: %v\n", err)
		}
	})
	if err != nil {
		fmt.Printf("Error scheduling market jobs: %v\n", err)
	}
}

func (cr *CronJob) ScheduleDailyMarketCheck(collection *mongo.Database) {
	store.MarketStatus(collection.Collection("market-status"))
	_, err := cr.c.AddFunc("2 11 * * 0-4", func() {
		log.Print("started")
		store.MarketStatus(collection.Collection("market-status"))
	})
	if err != nil {
		fmt.Print("error scheduling utils init function")
	}

	_, err = cr.c.AddFunc("3 15 * * 0-4", func() {
		fmt.Println("Cron job with ID 3 is scheduled.")
		store.MarketStatus(collection.Collection("market-status"))
	})

	if err != nil {
		fmt.Print("error scheduling utils ismarketopen function")
	}
}

func (cr *CronJob) ScheduleDailyMarketData(mongodatabase *mongo.Database) {
	store.StoreOrUpdateMarketData(mongodatabase.Collection("market-data"))
	_, err := cr.c.AddFunc("5-59/1 11 * * 0-4", func() {
		fmt.Println("Cron job with ID 4 is scheduled.")
		// Add your task here
		store.StoreOrUpdateMarketData(mongodatabase.Collection("market-data"))
	})

	if err != nil {
		fmt.Printf("Error scheduling market data jobs: %v\n", err)
	}

	_, err = cr.c.AddFunc("0-59/1 12-14 * * 0-4", func() {
		fmt.Println("Cron job with ID 5 is scheduled.")
		// Add your task here
		store.StoreOrUpdateMarketData(mongodatabase.Collection("market-data"))
	})

	if err != nil {
		fmt.Printf("Error scheduling market data jobs: %v\n", err)
	}

	_, err = cr.c.AddFunc("0-1/1 15 * * 0-4", func() {
		fmt.Println("Cron job with ID 6 is scheduled.")
		// Add your task here
		store.StoreOrUpdateMarketData(mongodatabase.Collection("market-data"))
	})

	if err != nil {
		fmt.Printf("Error scheduling market data jobs: %v\n", err)
	}
}

func (cr *CronJob) ScheduleIPOAndFPOData(mongodatabase *mongo.Database) {
	store.StoreIpoandFpoData(mongodatabase.Collection("ipo-fpo"))
	_, err := cr.c.AddFunc("0 9-18/2 * * *", func() {
		fmt.Println("Cron job with ID 7 is scheduled.")
		// Add your task here
		store.StoreIpoandFpoData(mongodatabase.Collection("ipo-fpo"))
	})

	if err != nil {
		fmt.Printf("Error scheduling IPO and FPO data jobs: %v\n", err)
	}
}

func (cr *CronJob) InitScheduler() {
	mongoDatabase := cr.MongoClient.Database("nepsedata")

	//schedule marketsummary functions
	// cr.ScheduleDailyMarketData(mongoDatabase)
	cr.ScheduleDailyMarketJobs(mongoDatabase)
	cr.ScheduleDailyMarketCheck(mongoDatabase)
	cr.ScheduleIPOAndFPOData(mongoDatabase)
	log.Print("scheduled")

}
