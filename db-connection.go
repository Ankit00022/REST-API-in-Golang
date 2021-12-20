package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:mysql@22@tcp(127.0.0.1:3306)/db?parseTime=true"
var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Print(err.Error())
		panic("connection failed:(")
	}
	Database.AutoMigrate(&Employee{})

}
