package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"ginlite/helpers"
	"ginlite/middlewares"
)

func main() {
	//init gin and global context
	r := gin.Default()

	//open db
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ginlite?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("mysql", "root:123456@/ginlint?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	r.Use(middlewares.SetContext("_self_context__db", db))
	r.Use(middlewares.GlobalAuth())

	//init helpers
	helpers.InitRand()

	RouteRegister(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
