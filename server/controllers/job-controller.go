package controllers

import (
	"net/http"

	"github.com/ivan-avalos/linkbucket-go/server/database"
	"github.com/ivan-avalos/linkbucket-go/server/utils"
	"github.com/labstack/echo/v4"
)

func GetFinishedJobs(c echo.Context) error {
	userID := utils.GetJWTUserID(c)
	jobs, err := database.GetFinishedJobs(userID)
	if err != nil {
		return utils.ProcessError(err)
	}
	responseJobs := make([]database.ResponseJob, len(jobs))
	for i, j := range jobs {
		responseJobs[i] = j.GetResponseJob()
	}
	return c.JSON(http.StatusOK, utils.BaseResponse(http.StatusOK, responseJobs))
}
