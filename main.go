package main

import (
	"github.com/difoul/go-rest-app/hello"
	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run() // listen and serve on 0.0.0.0:8080
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/v1")
	hello.RegisterHello(v1)
	return router
}
