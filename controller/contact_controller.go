package controller

import (
	"example/back/config"
	"example/back/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context){
	contacts := []model.Contact{}
	config.DB.Find(&contacts)
	c.JSON(http.StatusOK, &contacts)
}

func Create(c *gin.Context) {
	var contact model.Contact
	c.BindJSON(&contact)
	config.DB.Create(&contact)
	c.JSON(http.StatusOK, &contact)
}

func GetOne(c *gin.Context){
	var contact model.Contact
	config.DB.Where("id = ?", c.Param("id")).First(&contact)
	c.JSON(http.StatusOK, &contact)
}

func Update(c *gin.Context){
	var contact model.Contact
	config.DB.Where("id = ?", c.Param("id")).First(&contact)
	c.BindJSON(&contact)
	config.DB.Save(&contact)
	c.JSON(http.StatusOK,gin.H{"message": "Update contact"})
}

func DeleteOne(c *gin.Context){
	var contact model.Contact
	config.DB.Where("id = ?", c.Param("id")).Delete(&contact)
	c.JSON(http.StatusOK,gin.H{"message": "Delete contact"})
}

func DeleteAll(c *gin.Context){
	config.DB.Exec("DELETE FROM contacts")
	c.JSON(http.StatusOK,gin.H{"message": "Delete all contacts"})
}

func Create10Row(c *gin.Context) {
	contacts := []*model.Contact{
		&model.Contact{Name: "John Doe", Number: "1234567890"},
		&model.Contact{Name: "Jane Smith", Number: "0987654321"},
		&model.Contact{Name: "Michael Brown", Number: "9876543210"},
		&model.Contact{Name: "Emily Johnson", Number: "0123456789"},
		&model.Contact{Name: "Christopher Davis", Number: "6789012345"},
		&model.Contact{Name: "Jessica Wilson", Number: "5432109876"},
		&model.Contact{Name: "Matthew Miller", Number: "7890123456"},
		&model.Contact{Name: "Olivia Anderson", Number: "4567890123"},
		&model.Contact{Name: "Andrew Taylor", Number: "2345678901"},
		&model.Contact{Name: "Isabella Thomas", Number: "9012345678"},
	}
	result := config.DB.Create(contacts)
	if result.Error != nil {
		panic(result.Error)
	}
	c.JSON(http.StatusOK,gin.H{"message": "Add 10 rows"})
}