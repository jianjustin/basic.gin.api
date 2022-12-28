package tool

import (
	"basic.gin.api/goods/entity"
	"basic.gin.api/middleware/model"
)

func RegisterTables() {
	model.DB.AutoMigrate(entity.Good{}) //注册商品表
}
