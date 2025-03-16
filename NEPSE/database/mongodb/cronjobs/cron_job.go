package cronjobs

import (
	"nepseserver/database/mongodb"
	applog "nepseserver/log"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
)

type CronJob struct {
	MongoClient *mongo.Client
	c           *cron.Cron
}

func InitMongo() *mongo.Client {
	return mongodb.Init()
}

func InitCronJobs(c *cron.Cron) *cron.Cron {
	return c

}

func NewCronJob(c *cron.Cron) *CronJob {

	applog.Log(applog.DEBUG, "Cron Job Starting")
	return &CronJob{
		MongoClient: InitMongo(),
		c:           InitCronJobs(c),
	}
}
