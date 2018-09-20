package main

import (
	"ginlite/controllers"
	"github.com/gin-gonic/gin"
)

func RouteRegister(r *gin.Engine) {

    v1 := r.Group("/api/v1")
    {
        v1.GET("/ping", (&controllers.TestController{}).Ping)
        v1.POST("/users", (&controllers.UserController{}).Create)
        v1.POST("/users.delete", (&controllers.UserController{}).Delete)
    }


}
