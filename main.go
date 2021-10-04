package main

import {
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
}

var db gorm.db
var err error

func main(){
	// fmt.Println("Hallo kita buat crud golang")
	db, err = gorm.Open("mysql", "root:@/go_rest_api_crud?charset=utf8&parseTime=True")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}
}