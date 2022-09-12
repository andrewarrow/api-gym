package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	router := gin.Default()

	router.GET("/", Welcome)
	return router
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": ""})
}
