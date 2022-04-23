package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/DockerizedGolangTemplate/internal/health"
	"github.com/sirupsen/logrus"
)

const HealthGroup = "/health"
const helloEndpoint = "/hello"

func BuildHelloEndpoint(logger *logrus.Logger, handlers ...gin.HandlerFunc) Endpoint {
	return Endpoint{
		URL:      helloEndpoint,
		Handlers: append(handlers, health.GetHelloHandler(logger)),
	}
}
