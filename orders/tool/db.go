package tool

import (
	"basic.gin.api/middleware/model"
	"basic.gin.api/orders/entity"
)

func RegisterTables() {
	model.DB.AutoMigrate(entity.Order{}) //注册订单表
}
