package controllers

import (
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    "ginlite/models"
)

type UserController struct {
}


type Product struct {
    gorm.Model
    Code string
    Price uint
}

//type User struct {
    //gorm.Model
    //UserName string
    //NickName string
    //AccountID string
//}


func (ctrl *UserController) Create(c *gin.Context) {
    //var user models.User

    db := c.MustGet("db").(*gorm.DB)

    fmt.Println("3222222222", c.PostForm("username"))
    fmt.Println("4222222222", c.PostForm("nickname"))

    db.Create(&models.User{
        Username: c.PostForm("username"),
        Nickname: c.PostForm("nickname"),
        Male: true,
        AccountID: "12345",
        Desc: "kkk",
        Detail: "hhh"})

    //c.JSON(200, user)
    c.JSON(200, gin.H{"status": "ok"})
}


func (ctrl *UserController) Delete(c *gin.Context) {
    //var user models.User

    db := c.MustGet("db").(*gorm.DB)

    fmt.Println("3222222222", c.PostForm("username"))
    fmt.Println("4222222222", c.PostForm("nickname"))

    db.Where("account_id=?", "12345").Delete(&models.User{})

    //c.JSON(200, user)
    c.JSON(200, gin.H{"status": "ok"})
}




