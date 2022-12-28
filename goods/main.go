package main

import (
	"basic.gin.api/goods/entity"
	"basic.gin.api/goods/tool"
	"basic.gin.api/middleware"
	"basic.gin.api/middleware/model"
	"github.com/gin-gonic/gin"
)

func main() {
	model.DB = middleware.BuildGormMysql()
	tool.RegisterTables()

	/*
		//create a good
		good := &entity.Good{
			GoodName: "good1",
			Price:    1.23,
		}
		model.DB.Create(&good)
	*/

	good := &entity.Good{}
	model.DB.First(&good)

	r := gin.Default()
	r.GET("/good", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"good": good,
		})
	})
	r.Run()
}
