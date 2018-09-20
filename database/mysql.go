package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var mysqlDB *gorm.DB

func Init() (*gorm.DB, error) {
    var err error

	//open db
	mysqlDB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/ginlite?charset=utf8&parseTime=True&loc=Local")

    return mysqlDB, err
}

func Get() *gorm.DB {
    return mysqlDB
}
