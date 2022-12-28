package tool

import (
	"basic.gin.api/middleware/model"
	"basic.gin.api/users/entity"
)

func RegisterTables() {
	model.DB.AutoMigrate(entity.SysUser{}) //注册用户表
}
