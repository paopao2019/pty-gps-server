package main

import (
	"PTY_GPS/api"
	"PTY_GPS/global"
	"PTY_GPS/initialize"
	"PTY_GPS/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())

	// 数据库mysql
	global.DB = initialize.Gorm()

	r.GET("/distance/all", api.GetDistanceList)
	r.GET("/distance/byName", api.GetDistanceByName)

	station := r.Group("/station")
	{
		station.POST("/addStation", api.AddStation)
		station.DELETE("/deleteStation", api.DeleteStation)
		station.POST("/getStationList", api.GetStationList)
		station.PUT("/updateStation", api.UpdateStation)
		station.GET("/calculate", api.CalculateStation)
	}
	r.Run(":9991")

}
