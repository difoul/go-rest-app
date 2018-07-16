package quotes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterQuotes(router *gin.RouterGroup) {
	router.GET("/random", GetRandomQuote)
	router.GET("/quote/:id", GetQuoteById)
}

func GetRandomQuote(c *gin.Context) {
	c.JSON(http.StatusOK, RandomQuote())
}

func GetQuoteById(c *gin.Context) {
	id_s := c.Param("id")
	id, err := strconv.Atoi(id_s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	q, err := QuoteById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, q)
}
