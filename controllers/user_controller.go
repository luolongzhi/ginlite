package controllers

import (
	"fmt"
    "strings"
	"ginlite/models"
	"github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type UserController struct {
}

type Input1 struct {
    Username string `form:"username" json:"username"`
    Nickname string `form:"nickname" json:"nickname"`
}

func (ctrl *UserController) Create(c *gin.Context) {
	//var user models.User
	var err error

	db := c.MustGet("db").(*gorm.DB)

	fmt.Println("3222222222", c.Param("username"))
	fmt.Println("4222222222", c.Param("nickname"))

	var input1 Input1

	contentType := c.Request.Header.Get("Content-Type")
    fmt.Println("contentType: ------------", contentType)
    fmt.Println("mmmmmmmm: ", strings.Contains(contentType, "multipart/form-data"))

    if (strings.Contains(contentType, "application/json")) {
        err = c.BindJSON(&input1)
    } else if (strings.Contains(contentType, "application/x-www-form-urlencoded")) {
		err = c.BindWith(&input1, binding.Form)
    } else if (strings.Contains(contentType, "multipart/form-data")) {
		err = c.BindWith(&input1, binding.Form)
    }


	if err != nil {
		fmt.Println(err)
	}

	//c.Bind(&input1)
	fmt.Println("input: ****************: ", input1)
	fmt.Println("input11: ****************: ", input1.Username, input1.Nickname)

	user := &models.User{
		Username:  c.PostForm("username"),
		Nickname:  c.PostForm("nickname"),
		Male:      true,
		AccountID: "123456789",
		Desc:      "kkk",
		Detail:    "hhh"}
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("eeeeeeeeeeee: ", err)
	}
	fmt.Println("kkkkkk: ", user.ID)
	fmt.Println("kkkkkk: ", *user)

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
