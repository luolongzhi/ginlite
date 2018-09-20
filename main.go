package main

import (
	"github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger" // gin-swagger middleware
    "github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files

    _ "ginlite/docs"
    "ginlite/database"
	"ginlite/helpers"
	"ginlite/middlewares"
)

// @title SmartShop API
// @version v1 
// @description This is a sample server celler server.
// @host localhost:8080
// @BasePath /api
func main() {
	//init gin and global context
	r := gin.Default()

    db, err := database.Init()
    if err != nil {
		panic("failed to connect database")
    }
	defer db.Close()

	r.Use(middlewares.GlobalAuth())

	//init helpers
	helpers.InitRand()

	RouteRegister(r)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080
}
