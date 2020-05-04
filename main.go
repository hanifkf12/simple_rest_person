package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"hanif.com/coba/config"
	"hanif.com/coba/controllers"
	"os"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}
	defer db.Close()
	router := gin.Default()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}
	gin.SetMode(gin.ReleaseMode)
	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person/:id", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)
	router.Run(":" + port)
}
