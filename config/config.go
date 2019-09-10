package config

import (
	"../structs"
	"fmt"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB{
	//db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/cobagodb?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=person_test password=hanif sslmode=disable")

	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
