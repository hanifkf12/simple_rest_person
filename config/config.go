package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"hanif.com/coba/structs"
)

func DBInit() *gorm.DB {
	//db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/cobagodb?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=person_test password=hanif sslmode=disable")
	db, err := gorm.Open("postgres", "host=ec2-52-72-221-20.compute-1.amazonaws.com port=5432 user=kuyoqdykqsvido dbname=ddii60l9jo7pn3 password=6f26bd6800b32023ec392c5b4c554ae50ef23675d2958307f0f9135eb8a7e72a sslmode=disable")

	if err != nil {
		fmt.Print(err)
		panic("failed to connect to database")
	}
	db.AutoMigrate(structs.Person{})
	return db
}
