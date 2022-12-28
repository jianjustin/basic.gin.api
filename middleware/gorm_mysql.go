package middleware

import (
	"basic.gin.api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func BuildGormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         255,                                                                             // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                           // 根据版本自动配置
	}

	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		panic("create mysql connection error")
	}
	return db
}

func RegisterTables() {
	model.DB.AutoMigrate(model.SysUser{}) //表注册
}
