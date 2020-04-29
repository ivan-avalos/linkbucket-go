package jobs

import (
	"log"

	"github.com/ivan-avalos/linkbucket-go/server/database"
)

// AsyncOldImport starts async job for OldImport
func AsyncOldImport(job Job) {
	userID := job.UserID
	links := job.Params[0].([]database.OldImportLink)
	for i := len(links) - 1; i >= 0; i-- {
		oil := links[i]
		link := oil.GetLink()
		link.UserID = userID
		if err := link.Create(); err != nil {
			log.Printf("Error %e on OldImport for %d", err, userID)
			return
		}
		if err := link.LinkTags(oil.Tags); err != nil {
			log.Printf("Error %e on OldImport for %d", err, userID)
			return
		}
	}
}
