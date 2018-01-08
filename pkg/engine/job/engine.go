package job

import (
	"time"
	"crypto-price-fetcher/pkg/models"
	"crypto-price-fetcher/pkg/database"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

func StartEngine(config models.Config) () {
	for {
		fmt.Println("tick")
		jobs := database.FindStartedJobs()
		fmt.Println(jobs)
		for _, job := range jobs {
			if waitDurationSurpassed(job) {
				executeJob(job)
			}
		}
		time.Sleep(time.Millisecond * time.Duration(config.SleepTimeInMilliseconds))
	}
}

func waitDurationSurpassed(job models.Job) (bool) {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()-job.LastExecutionTime > job.WaitDuration
}

func executeJob(job models.Job) () {
	fmt.Println("Executing job \"" + job.Id.Hex() + "\"")
	database.UpdateJob(job.Id.Hex(), bson.M{"last-execution-time": time.Now().Unix()})
}
