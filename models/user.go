package models

import (
    "time"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)


type User struct {
    ID uint `gorm:"primary_key" json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time `sql:"index" json:"deleted_at"`

    Username string `json:"username"`
    Nickname string `gorm:"default:'我是小白龙'" json:"nickname"`
    Male bool `gorm:"default:true" json:"male"`
    AccountID string `json:"account_id"`
    Desc string `json:"desc"`
    Detail string `json:"detail"`

}

