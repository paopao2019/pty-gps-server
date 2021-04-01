package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDistanceList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
			"data": "你好",
	})
}


func GetDistanceByName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "ByName",
	})
}
