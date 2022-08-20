package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "gouser:gouser@tcp(localhost:3306)/userdb?parseTime=true"
var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&Employee{})
}
