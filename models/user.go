package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    gorm.Model

    Username string `json:"username"`
    Nickname string `gorm:"default:'我是小白龙'" json:"nickname"`
    Male bool `gorm:"default:true" json:"male"`
    AccountID string `json:"account_id"`
    Desc string `json:"desc"`
    Detail string `json:"detail"`
}

