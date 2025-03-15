package cronjobs

import (
	"fmt"
	"nepseserver/database/mongodb"

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

	fmt.Print("Cron Job Starting")
	return &CronJob{
		MongoClient: InitMongo(),
		c:           InitCronJobs(c),
	}
}
