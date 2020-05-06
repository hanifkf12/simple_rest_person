package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hanif.com/coba/structs"
	"net/http"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	token := c.Request.Header.Get("Authorization")
	fmt.Print("token nya", token)
	id := c.Param("id")
	//err := idb.DB.Where("id = ?", id).First(&person).Error
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": err.Error() + " with id " + id,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)
	idb.DB.Order("id").Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	err := c.ShouldBind(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}
	c.JSON(http.StatusCreated, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusOK, result)
		return
	}
	e := c.ShouldBind(&newPerson)
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil || e != nil {
		result = gin.H{
			"result": "update failed",
		}
		c.JSON(http.StatusOK, result)
		return
	} else {
		result = gin.H{
			"result": "successfully updated data",
			"data":   person,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	fmt.Println(person.First_Name, person.Last_Name)
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusOK, result)
		return
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
		c.JSON(http.StatusOK, result)
		return
	} else {
		result = gin.H{
			"result": "Data deleted successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}
