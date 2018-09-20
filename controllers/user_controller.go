package controllers

import (
	"fmt"
	//"strconv"
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"github.com/jinzhu/gorm"

    "ginlite/database"
	h "ginlite/helpers"
	m "ginlite/models"
)

type UserController struct {
}


type CreateParam struct {
	Username string `form:"username" json:"username" binding:"required"`
	Nickname string `form:"nickname" json:"nickname"`
	Male bool `form:"male" json:"male"`
}
// 创建用户 
// @Summary 创建用户 
// @Description 创建用户 
// @Accept  mpfd
// @Produce  json
// @Param   username formData string true  "用户名字段"
// @Param   nickname formData string false "用户昵称"
// @Param   male formData bool false "男性"
// @Success 200 {object} models.User helpers.ErrorBody
// @Failure 400 {object} helpers.ErrorBody "错误码"
// @Router /v1/users [post]
func (c *UserController) Create(ctx *gin.Context) {
	var param CreateParam
    var	err   error
    var user  m.User

	db := database.Get()
	if err = ctx.ShouldBind(&param); err != nil {
		h.ResError(ctx, fmt.Sprintf("%s", err))
		return
	}

	user = m.User{
		Username: param.Username,
		Nickname: param.Nickname,
		Male:     true,
		Desc:   "kkk",
		Detail: "hhh",
    }

	if err = db.Create(&user).Error; err != nil {
		h.ResError(ctx, fmt.Sprintf("%s", err))
		return
	}

	db.First(&user, user.ID)

	h.ResJson(ctx, user)
}




type UpdateParam struct {
	Username string `form:"username" json:"username"`
	Nickname string `form:"nickname" json:"nickname"`
    Male bool `form:"male" json:"male" binding:"exists"`
}
// 更新用户 
// @Summary 更新用户 
// @Description 更新用户
// @Accept  mpfd
// @Produce  json
// @Param   id path string true "用户ID"
// @Param   username formData string false  "用户名字段"
// @Param   nickname formData string false "用户昵称"
// @Param   male formData bool false "男性"
// @Success 200 {object} models.User helpers.ErrorBody
// @Failure 400 {object} helpers.ErrorBody "错误码"
// @Router /v1/users/{id} [put]
func (c *UserController) Update(ctx *gin.Context) {
	var param UpdateParam
    var	err   error
    var user  m.User

    id := ctx.Param("id")
	db := database.Get()
	if err = ctx.ShouldBind(&param); err != nil {
		h.ResError(ctx, fmt.Sprintf("%s", err))
		return
	}

    fmt.Println("id: ", id)
    fmt.Println("username: ", param.Nickname)
    fmt.Println("nickname: ", param.Nickname)
    fmt.Println("male: ", param.Male)

	db.First(&user, id)
	if err = db.Model(&user).Update(param).Error; err != nil {
		h.ResError(ctx, fmt.Sprintf("%s", err))
		return
	}

	db.First(&user, user.ID)

	h.ResJson(ctx, user)

}

func (ctrl *UserController) Delete(ctx *gin.Context) {
	//var user models.User

	db := database.Get()

	fmt.Println("3222222222", ctx.PostForm("username"))
	fmt.Println("4222222222", ctx.PostForm("nickname"))

	db.Where("account_id=?", "12345").Delete(&m.User{})

	ctx.JSON(200, gin.H{"status": "ok"})
}
