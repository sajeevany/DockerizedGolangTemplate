package pd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/DockerizedGolangTemplate/internal/endpoints"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	Group = "/pd"
	User  = "/users"

	//pagerduty API header args
	tokenValueFormat = "Token token=%s"

	//PD get Users url
	pdGetUsersURL = "https://api.pagerduty.com/users"
)

func BuildGetUsersEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) endpoints.Endpoint {
	return endpoints.Endpoint{
		URL:      User,
		Handlers: append(handlers, getUsersHandler(logger)),
	}
}

type Ping struct {
	Response string `json:"response" required:"true" description:"Server hello response" example:"hello"`
}

type UsersHeader struct {
	Token string `header:"token" binding:"required"`
}

func getUsersHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Sample logging message
		logger.Debug("Handling a pd request")

		//validate header requirements
		var header UsersHeader
		if err := ctx.ShouldBindHeader(&header); err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		// build GET request
		req, err := http.NewRequest(http.MethodGet, pdGetUsersURL, nil)
		if err != nil {
			logger.Debug("unable to setup http get request. err <%v>", pdGetUsersURL, err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		req.Header.Set("Content-Type", gin.MIMEJSON)
		req.Header.Set("Authorization", fmt.Sprintf(tokenValueFormat, header.Token))
		req.Header.Set("Accept", "application/vnd.pagerduty+json;version=2")

		//run request
		client := http.Client{
			Timeout: 2 * time.Second,
		}
		resp, err := client.Do(req)
		if err != nil {
			ctx.AbortWithError(http.StatusBadGateway, err)
			return
		}
		defer resp.Body.Close()

		//return body as received
		ctx.DataFromReader(resp.StatusCode, resp.ContentLength, resp.Header.Get("content-type"), resp.Body, nil)
	}

}
