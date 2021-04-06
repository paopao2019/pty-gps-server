package service

import (
	"PTY_GPS/global"
	"PTY_GPS/model"
	"PTY_GPS/uitls"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func AddStation(station model.PTYStation) (err error) {
	if !errors.Is(global.DB.Where("code = ? ", station.Code).First(&model.PTYStation{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同地点")
	}
	// Geo获取坐标地址
	err, location := uitls.GetPosition(station.Address, "上海")
	if err != nil {
		return err
	}
	station.Location = location
	err = global.DB.Create(&station).Error
	return err
}


// 删
func DeleteStation(id float64) (err error) {
	// 1. 判断不存在没有在其他 表中使用 虚拟机中 <- 待补充
	var station model.PTYStation
	// 2. 删除
	err = global.DB.Where("id = ?", id).Delete(&station).Error
	return err
}

// 改
// 通过唯一ID进行更新
func UpdateStation(station model.PTYStation) (err error) {
	// 通过唯一ID进行更新
	// Geo获取坐标地址
	err, location := uitls.GetPosition(station.Address, "上海")
	if err != nil {
		return err
	}
	station.Location = location
	err = global.DB.Where("id = ?", station.ID).First(&model.PTYStation{}).Updates(&station).Error
	return err
}

// 查
func GetStationInfoList(station model.PTYStation, info model.PageInfo, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.DB.Model(&model.PTYStation{})
	var stationList []model.PTYStation
	if station.Code != "" {
		db = db.Where("code = ?", station.Code)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, stationList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		err = db.Find(&stationList).Error
	}
	return err, stationList, total
}

// 计算
func CalculateStation(originalLocation string) (err error, list interface{}){
	var calcResult []model.StationCalcInfo
	var stationList []model.PTYStation
	db := global.DB.Model(&model.PTYStation{})
	db.Find(&stationList)
	fmt.Printf("葡萄园,stationList: %v", stationList)
	for _, v := range stationList {
		// 重组数据
		err, distance, duration := uitls.GetDistance(originalLocation, v.Location)
		err, tranTime := uitls.GetTranTime(originalLocation, v.Location)
		if err != nil {
			distance = "none"
			duration = "none"
			fmt.Printf("计算: %s 的距离出错: %v", v.Address, err.Error())
			continue
		}
		calcResult = append(calcResult, model.StationCalcInfo{
			PTYStation: v,
			Distance: distance,
			Duration: duration, // 驾车时间
			TranTime: tranTime, // 乘坐交通的最优时间 预估时间
		})
	}
	return  err, calcResult

}