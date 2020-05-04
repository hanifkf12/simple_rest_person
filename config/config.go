package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"hanif.com/coba/structs"
)

func DBInit() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/cobagodb?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=person_test password=hanif sslmode=disable")
	db, err := gorm.Open("postgres", "host=ec2-18-215-99-63.compute-1.amazonaws.com port=5432 user=dtyopzwmqpjlfb dbname=dfbeou4tihjq9l password=7dc57c7f445fe2e2170324592c98752343475abf3a1d72af57fbd67ee4fa14dc sslmode=disable")

	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	db.AutoMigrate(structs.Person{})
	return db
}
