package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"hanif.com/coba/config"
	"hanif.com/coba/structs"
)

func main() {
	db := config.DBInit()
	db.AutoMigrate(structs.Person{})
	fmt.Println("Migration Run")
}
