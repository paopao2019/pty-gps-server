package model

import "gorm.io/gorm"




type PTYStation struct {
	gorm.Model
	Code     string   `json:"code" gorm:"type:varchar(256);not null;comment:教学点代码"`
	Address  string `json:"address" gorm:"type:varchar(256);not null;comment:教学点地址"`
	Location string `json:"location" gorm:"type:varchar(256);not null;comment:坐标" `
}


type SearchStationParams struct {
	PTYStation
	PageInfo
	Desc     bool   `json:"desc"`
}

// 计算结果信息
type StationCalcInfo struct {
	PTYStation
	Distance string `json:"distance"`
	Duration string `json:"duration"`
}