package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sajeevany/DockerizedGolangTemplate/internal/endpoints"
)

const v1Api = "/api/vi"

func main() {
	router := setupRouter()
	setupV1Routes(router)
	router.Run()
}

func setupRouter() *gin.Engine {
	return gin.Default()
}

func setupV1Routes(rtr *gin.Engine) {
	addHealthEndpoints(rtr)
}

func addHealthEndpoints(rtr *gin.Engine) {
	v1 := rtr.Group(endpoints.HealthGroup)
	{
		hello := endpoints.BuildHelloEndpoint()
		v1.GET(hello.URL, hello.Handlers ...)
	}
}
