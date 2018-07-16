package quotes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterQuotes(router *gin.RouterGroup) {
	router.GET("/random", GetRandomQuote)
}

func GetRandomQuote(c *gin.Context) {
	c.JSON(http.StatusOK, RandomQuote())
}
