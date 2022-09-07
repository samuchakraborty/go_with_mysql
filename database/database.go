package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Database() *gorm.DB {

	// db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/onlineShopping")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// return db

	dsn := "root:@tcp(127.0.0.1:3306)/onlineShopping?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
