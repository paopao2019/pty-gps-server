package api

import (
	"PTY_GPS/model"
	"PTY_GPS/service"
	"PTY_GPS/uitls"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 增
func AddStation(c *gin.Context) {
	var station model.PTYStation
	_ = c.ShouldBindJSON(&station)
	if err := service.AddStation(station); err != nil {
		fmt.Printf("创建失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "创建失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "创建成功",
		})
	}
}


// 删
func DeleteStation(c *gin.Context) {
	var station model.GetById
	_ = c.ShouldBindJSON(&station)

	if err := service.DeleteStation(station.Id); err != nil {
		fmt.Printf("创建失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "删除失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "删除成功",
		})
	}
}


// 改
func UpdateStation(c *gin.Context) {
	var R model.PTYStation
	_ = c.ShouldBindJSON(&R)

	if err := service.UpdateStation(R); err != nil {
		fmt.Printf("更新失败: %v", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "更新失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "更新成功",
		})
	}
}

// 查
func GetStationList(c *gin.Context) {
	var pageInfo model.SearchStationParams
	_ = c.ShouldBind(&pageInfo)
	fmt.Println(pageInfo.Code)
	if err, list, total := service.GetStationInfoList(pageInfo.PTYStation, pageInfo.PageInfo, pageInfo.Desc); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "获取失败",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "获取成功",
			"data": map[string]interface{} {
				"list": list,
				"total": total,
				"page": pageInfo.Page,
				"pageSize": pageInfo.PageSize,
			},
		})
	}
}

// calculateStation
func CalculateStation(c *gin.Context) {
	originalAddr := c.Query("original")
	log.Print("请求地址是:",originalAddr)
	_, originalLocation := uitls.GetPosition(originalAddr, "上海")
	fmt.Printf("originalLocation is %s", originalLocation)


	// 查找地址
	if err, list := service.CalculateStation(originalLocation); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg": "获取失败",
		})

	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg": "获取成功",
			"data": map[string]interface{} {
				"list": list,
			},
		})
	}
}