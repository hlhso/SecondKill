package database

import (
	"SecondKill/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var dbInstance *gorm.DB

// 单例模式
func GetMysqlInstance() *gorm.DB {
	if dbInstance == nil {
		dbInstance = initDB()
	}

	return dbInstance
}

func initDB() *gorm.DB {
	db, err := gorm.Open(config.DataBaseSetting.Type, config.DataBaseSetting.SqlDSN)
	// 设置为true,数据操作日志可以数据在控制台
	db.LogMode(true)
	// 默认情况下创建表明结构为复数
	db.SingularTable(true)

	// Error
	if err != nil {
		config.AppSetting.Logger.Error("连接数据库不成功", err, config.DataBaseSetting)
	} else {
		config.AppSetting.Logger.Info("数据库 链接成功\n")
	}

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(config.DataBaseSetting.MaxIdle)
	//打开
	db.DB().SetMaxOpenConns(config.DataBaseSetting.MaxOpen)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * config.DataBaseSetting.IdleTimeout)

	return db
}

