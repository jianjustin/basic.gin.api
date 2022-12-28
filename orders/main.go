package main

import (
	"basic.gin.api/middleware"
	"basic.gin.api/middleware/model"
	"basic.gin.api/orders/entity"
	"basic.gin.api/orders/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	model.DB = middleware.BuildGormMysql()
	tool.RegisterTables()

	order := &entity.Order{
		GoodId:    1,
		GoodCount: 10,
		Sum:       10,
	}
	model.DB.Create(&order)

	model.DB.First(&order)

	r := gin.Default()
	r.GET("/order", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user": order,
		})
	})
	r.Run()
}
