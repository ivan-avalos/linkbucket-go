package jobs

import (
	"encoding/json"
	"strings"

	"github.com/ivan-avalos/linkbucket-go/server/database"
)

// AsyncOldImport starts async job for OldImport
func AsyncOldImport(job database.Job) {
	userID := job.UserID
	var params [][]database.OldImportLink
	err := json.Unmarshal([]byte(job.FnParams), &params)
	links := params[0]
	if err != nil {
		job.Error = err.Error()
		job.LogError()
		job.Update()
		return
	}
	for i := len(links) - 1; i >= 0; i-- {
		oil := links[i]
		link := oil.GetLink()
		link.UserID = userID
		ok, err := link.IsUnique()
		if err != nil {
			job.Error = err.Error()
			job.LogError()
			job.Update()
			return
		}
		if !ok {
			continue
		}
		if err := link.Create(); err != nil {
			job.Error = err.Error()
			job.LogError()
			job.Update()
			return
		}
		trimmedTags := make([]string, len(oil.Tags))
		for i, t := range oil.Tags {
			trimmedTags[i] = strings.Trim(t, ",")
		}
		if err := link.LinkTags(trimmedTags); err != nil {
			job.Error = err.Error()
			job.LogError()
			job.Update()
			return
		}
	}
	job.Done = true
	job.Update()
}
