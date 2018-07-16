package hello

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterHello(router *gin.RouterGroup){
  router.GET("/hello", SayHello)
}

func SayHello(c *gin.Context){
  c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
