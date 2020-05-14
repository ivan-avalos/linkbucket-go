package database

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
)

type (
	// Job represents an async job
	Job struct {
		gorm.Model
		UserID    uint
		Name      string
		FnName    string
		FnParams  string `gorm:"type:longtext"`
		RetryLeft uint
		Done      bool
		Error     string
		Params    []interface{} `gorm:"-"`
	}

	// ResponseJob represents a Job for response
	ResponseJob struct {
		ID    uint        `json:"id"`
		Name  string      `json:"name"`
		Done  bool        `json:"done"`
		Error interface{} `json:"error"`
	}

	// CallMap represents map to resolve function calls
	CallMap map[string]func(Job)
)

// GetResponseJob returns ResponseJob from Job
func (job *Job) GetResponseJob() ResponseJob {
	var error interface{} = nil
	if job.Error != "" {
		error = job.Error
	}
	return ResponseJob{
		ID:    job.ID,
		Name:  job.Name,
		Done:  job.Done,
		Error: error,
	}
}

// Create inserts Job into DB
func (job *Job) Create() error {
	retry, err := strconv.Atoi(os.Getenv("QUEUE_RETRY"))
	if err != nil {
		return err
	}
	job.RetryLeft = uint(retry)
	params, err := json.Marshal(job.Params)
	if err != nil {
		return err
	}
	job.FnParams = string(params)
	return DB().Create(job).Error
}

// Start dispatches Job by calling mapped function
func (job *Job) Start(cm CallMap) {
	job.RetryLeft--
	job.Done = true
	job.Update()

	json.Unmarshal([]byte(job.FnParams), &job.Params)

	log.Printf("Job %s for user %d started", job.Name, job.UserID)
	cm[job.FnName](*job)
	log.Printf("Job %s for user %d finished", job.Name, job.UserID)
}

// LogError prints error message for Job
func (job *Job) LogError() {
	log.Printf("Job %s for user %d error: %s", job.Name, job.UserID, job.Error)
}

// GetUnfinishedJobs returns last 50 unfinished/failed jobs (used for processing)
func GetUnfinishedJobs() ([]Job, error) {
	var jobs []Job
	err := DB().Limit(50).
		Where("retry_left > 0").
		Where("done = ? OR error IS NOT NULL", true).
		Find(&jobs).Error
	return jobs, err
}

// GetFinishedJobs returns last 10 finished/failed jobs for user (used for polling)
func GetFinishedJobs(userID uint) ([]Job, error) {
	var jobs []Job
	err := DB().Limit(10).Order("id desc").
		Where("user_id = ?", userID).
		Find(&jobs).Error
	return jobs, err
}

// Update modifies Job in DB
func (job *Job) Update() error {
	return DB().Save(job).Error
}

// Delete removes Job from DB
func (job *Job) Delete() error {
	return DB().Delete(job).Error
}
