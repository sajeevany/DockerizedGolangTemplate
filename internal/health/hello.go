package health

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/DockerizedGolangTemplate/internal/endpoints"
	"github.com/sirupsen/logrus"
	"net/http"
)

const Group = "/health"
const helloEndpoint = "/hello"

func BuildGetHelloEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) endpoints.Endpoint {
	return endpoints.Endpoint{
		URL:      helloEndpoint,
		Handlers: append(handlers, getHelloHandler(logger)),
	}
}

type Ping struct {
	Response string `json:"response" required:"true" description:"Server hello response" example:"hello"`
}

func getHelloHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Sample logging message
		logger.Debug("Handling a hello request")

		//Set response
		ctx.JSON(http.StatusOK, Ping{Response: "hello"})
	}

}
