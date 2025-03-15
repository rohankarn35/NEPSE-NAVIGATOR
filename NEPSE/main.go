package main

import (
	"log"
	"nepseserver/constants"
	"nepseserver/database/mongodb"
	"nepseserver/database/mongodb/cronjobs"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	constants.InitConstant()

	loc := time.FixedZone("NPT", 5*60*60+45*60) // NPT is UTC+5:45
	time.Local = loc

	c := cron.New(cron.WithLocation(loc))

	cron := cronjobs.NewCronJob(c)
	cron.InitScheduler()
	mongoClient := mongodb.Init()
	if mongoClient == nil {
		log.Fatal("Failed to initialize MongoDB client")
	}
	c.Start()
	select {}
}
