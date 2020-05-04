package structs

import (
	"github.com/jinzhu/gorm"
)

type Person struct {
	gorm.Model
	First_Name string `form:"first_name" json:"first_name"`
	Last_Name  string `form:"last_name" json:"last_name"`
}
