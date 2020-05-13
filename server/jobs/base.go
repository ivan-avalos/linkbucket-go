package jobs

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ivan-avalos/linkbucket-go/server/database"
	"github.com/ivan-avalos/linkbucket-go/server/utils"
)

var jobChan chan database.Job

var callMap database.CallMap = database.CallMap{
	"jobs.AsyncOldImport": AsyncOldImport,
}

func loadFromDB() {
	jobs, err := database.GetUnfinishedJobs()
	if err != nil {
		log.Println("Could not load jobs from DB")
		log.Println(err)
	} else {
		for _, j := range jobs {
			Enqueue(j)
		}
	}
}

// Enqueue adds job to queue
func Enqueue(job database.Job) error {
	select {
	case jobChan <- job:
		return nil
	default:
		return utils.BaseError(
			http.StatusServiceUnavailable,
			"Max capacity reached",
			"max_capacity_error",
			nil,
		)
	}
}

// Init starts worker pool with queue
func Init() {
	queueSize, err := strconv.Atoi(os.Getenv("QUEUE_LIMIT"))
	if err != nil {
		log.Println("QUEUE_LIMIT is invalid or null")
	}
	queueNumber, err := strconv.Atoi(os.Getenv("QUEUE_NUMBER"))
	if err != nil {
		log.Println("QUEUE_NUMBER is invalid or null")
	}
	jobChan = make(chan database.Job, queueSize)
	loadFromDB()
	for i := 1; i < queueNumber; i++ {
		go func() {
			for j := range jobChan {
				j.Start(callMap)
			}
		}()
		log.Printf("Started worker %d", i)
	}
}
