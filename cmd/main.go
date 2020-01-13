package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/DockerizedGolangTemplate/internal/endpoints"
	"github.com/sajeevany/DockerizedGolangTemplate/internal/logging"
	"github.com/sirupsen/logrus"
)

const v1Api = "/api/v1"

func main() {

	//Create a universal logger
	logger := logging.Init()

	//Initialize router
	router := setupRouter(logger)

	//Setup routes
	setupV1Routes(router)

	//Use default route of 8080. TODO add config reader to get the port
	err := router.Run(":8080")
	if err != nil {
		logger.Error("An error occurred when starting the router. <%v>", err)
	}

}

//setupRouter - Create the router and set middleware
func setupRouter(logger *logrus.Logger) *gin.Engine {

	engine := gin.New()

	//Add middleware
	engine.Use(logging.SetCtxLogger(logger))
	engine.Use(logging.LogRequest(logger))
	engine.Use(gin.Recovery())

	return engine
}

func setupV1Routes(rtr *gin.Engine) {
	addHealthEndpoints(rtr)
}

func addHealthEndpoints(rtr *gin.Engine) {
	v1 := rtr.Group(fmt.Sprintf("%s%s", v1Api, endpoints.HealthGroup))
	{
		hello := endpoints.BuildHelloEndpoint()
		v1.GET(hello.URL, hello.Handlers...)
	}
}
