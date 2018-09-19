package main

import (
	"ginlite/controllers"
	"github.com/gin-gonic/gin"
)

func RouteRegister(r *gin.Engine) {
	r.GET("/ping", (&controllers.TestController{}).Ping)
	r.POST("/users", (&controllers.UserController{}).Create)
	r.POST("/users.delete", (&controllers.UserController{}).Delete)
}
