package main

import (
	"basic.gin.api/middleware"
	"basic.gin.api/model"
)

func main() {
	model.DB = middleware.BuildGormMysql()
	middleware.RegisterTables()

	user := &model.SysUser{
		UserName: "admin1",
		Password: "admin1",
	}

	model.DB.Create(&user)
}
