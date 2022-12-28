package main

import (
	"basic.gin.api/middleware"
	"basic.gin.api/model"
)

func main() {
	model.DB = middleware.BuildGormMysql()
}
