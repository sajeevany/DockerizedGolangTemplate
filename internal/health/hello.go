package health

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Ping struct {
	Response string `json:"response" required:"true" description:"Server hello response" example:"hello"`
}

func GetHelloHandler(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Sample logging message
		logger.Debug("Handling a hello request")

		//Set response
		ctx.JSON(http.StatusOK, Ping{Response: "hello"})
	}

}
