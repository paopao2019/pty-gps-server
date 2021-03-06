package initialize

import (
	"PTY_GPS/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	// 数据库mysql
	dsn := "root:p@ss1234@anji@tcp(10.108.26.60:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Mysql数据库无法连接, error: %v", err)
	} else {
		// 设置参数
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	}

	// 自动迁移 会自动和数据库对应 创建数据库表  如果表存在会自动重新更新数据库如果Model有更改
	db.Debug().AutoMigrate(&model.PTYStation{})   // 葡萄园教学点表
	return db
}

