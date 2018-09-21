package main

import (
	"ginlite/controllers"
	"github.com/gin-gonic/gin"
	"ginlite/middlewares"
)

func RouteRegister(r *gin.Engine) {

    r.GET("/ping", (&controllers.TestController{}).Ping)

    v1 := r.Group("/api/v1")
    {
        apiAccount := v1.Group("/")
        {
            apiAccount.Use(middlewares.TokenAuth())

            apiAccount.POST("/users", (&controllers.UserController{}).Create)
        }

        apiT := v1.Group("/")
        {
            apiT.Use(middlewares.AppAuth())
            apiT.POST("/users/:id", (&controllers.UserController{}).Update)
        }

        v1.DELETE("/users.delete", (&controllers.UserController{}).Delete)
    }


}
