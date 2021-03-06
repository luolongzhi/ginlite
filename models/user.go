package models

import (
	"fmt"
    //"reflect"
	h "ginlite/helpers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strconv"
	"time"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`

	Username  string `json:"username"`
	Nickname  string `gorm:"default:'我是小白龙'" json:"nickname"`
	Male      bool   `gorm:"default:true" json:"male"`
	AccountID string `json:"account_id"`
	Desc      string `json:"desc"`
	Detail    string `json:"detail"`

    Password string `json:"-"`
	Avatar   string
    Age      int    `json:"age"`
}

type RetUser struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`

	Username  string `json:"username"`
	Nickname  string `gorm:"default:'我是小白龙'" json:"nickname"`
	Male      bool   `gorm:"default:true" json:"male"`
	AccountID string `json:"account_id"`
	Desc      string `json:"desc"`
	Detail    string `json:"detail"`

    Avatar   string `json:"avatar"`
}

func (u *User) BeforeCreate() (err error) {
	u.AccountID = strconv.Itoa(h.RandIntRange(1000000000, 9999999999))
	u.Password = h.HashPassword(u.AccountID)

	fmt.Println(h.GenerateJwt("dev", u.ID, u.AccountID))
	return
}

//func (u *User) ToJson() {
    //t := reflect.TypeOf(*u)
    //v := reflect.ValueOf(*u)

    //fmt.Println("iiiiiiiiii----------%v", t.NumField)

    //for k := 0; k < t.NumField(); k++ {
        //fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
        //result = append(result, t.Field(i).Name)
    //}

    //result := make([]string,0,fieldNum)
    //for i:= 0;i<fieldNum;i++ {
    //}
    //return resultgithub

//}

