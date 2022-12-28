package main

import (
	"basic.gin.api/middleware"
	"basic.gin.api/middleware/model"
	"basic.gin.api/users/entity"
	"basic.gin.api/users/tool"
	"github.com/gin-gonic/gin"
)

func main() {
	model.DB = middleware.BuildGormMysql()
	tool.RegisterTables()

	user := &entity.SysUser{}
	model.DB.First(&user)

	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"user": user,
		})
	})
	r.Run()
}
