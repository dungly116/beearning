package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	var connectionString = "root:Bebibungbu1994!@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"
	d, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
