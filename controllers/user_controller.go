package controllers

import (
	"fmt"
	//"strconv"
	h "ginlite/helpers"
	m "ginlite/models"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
)

type UserController struct {
}


type CreateParam struct {
	Username string `form:"username" json:"username" binding:"required"`
	Nickname string `form:"nickname" json:"nickname"`
}

// 创建用户 
// @Summary 创建用户 
// @Description 创建用户 
// @Accept  mpfd
// @Produce  json
// @Param   username formData string true  "用户名字段"
// @Param   nickname formData string false "用户昵称"
// @Success 200 {object} models.User helpers.ErrorBody
// @Failure 400 {object} helpers.ErrorBody "错误码"
// @Router /v1/users [post]
func (ctrl *UserController) Create(c *gin.Context) {
	var (
		param CreateParam
		err   error
	)

	db := c.MustGet("_self_context__db").(*gorm.DB)
	if err = c.ShouldBind(&param); err != nil {
		h.ResError(c, fmt.Sprintf("%s", err))
		return
	}

	user := &m.User{
		Username: param.Username,
		Nickname: param.Nickname,
		Male:     true,
		//AccountID: strconv.Itoa(h.RandIntRange(10000000, 99999999)),
		Desc:   "kkk",
		Detail: "hhh"}
	if err = db.Create(user).Error; err != nil {
		h.ResError(c, fmt.Sprintf("%s", err))
		return
	}

	var retUser m.User
	db.First(&retUser, user.ID)

	h.ResJson(c, retUser)
}

func (ctrl *UserController) Delete(c *gin.Context) {
	//var user models.User

	db := c.MustGet("db").(*gorm.DB)

	fmt.Println("3222222222", c.PostForm("username"))
	fmt.Println("4222222222", c.PostForm("nickname"))

	db.Where("account_id=?", "12345").Delete(&m.User{})

	//c.JSON(200, user)
	c.JSON(200, gin.H{"status": "ok"})
}
