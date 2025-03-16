package main

import (
	"nepseserver/constants"
	"nepseserver/database/mongodb"
	"nepseserver/database/mongodb/cronjobs"
	applog "nepseserver/log"
	"time"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {

	err := applog.InitLogger("app.log", applog.INFO) // Set minimum level to INFO
	if err != nil {
		applog.Log(applog.ERROR, "Failed to initialize logger: %v", err)
		return
	}
	defer applog.CloseLogger()

	err = godotenv.Load()
	if err != nil {
		applog.Log(applog.ERROR, "Error loading .env file: %v", err)
		return
	}

	constants.InitConstant()

	loc := time.FixedZone("NPT", 5*60*60+45*60) // NPT is UTC+5:45
	time.Local = loc

	c := cron.New(cron.WithLocation(loc))

	cron := cronjobs.NewCronJob(c)
	cron.InitScheduler()
	mongoClient := mongodb.Init()
	if mongoClient == nil {
		applog.Log(applog.ERROR, "Failed to initialize MongoDB client")
		return
	}
	c.Start()
	select {}
}
