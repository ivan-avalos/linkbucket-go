package jobs

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/ivan-avalos/linkbucket-go/server/utils"
)

// Job represents async job
type Job struct {
	UserID   uint
	Name     string
	Function func(Job)
	Params   []interface{}
}

func (job *Job) start() {
	log.Printf("Job %s for user %d started", job.Name, job.UserID)
	job.Function(*job)
	log.Printf("Job %s for user %d finished", job.Name, job.UserID)
}

var jobChan chan Job

func worker() {
	for job := range jobChan {
		job.start()
	}
}

// Init initialises queue worker
func Init() {
	limit, err := strconv.Atoi(os.Getenv("QUEUE_LIMIT"))
	if err != nil {
		log.Panic("QUEUE_LIMIT is invalid or null")
	}
	jobChan = make(chan Job, limit)
	worker()
}

// Enqueue adds job to queue if there is capacity
func Enqueue(job Job) error {
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
