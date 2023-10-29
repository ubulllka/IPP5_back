package routes

import (
	"example/back/controller"
	"github.com/gin-gonic/gin"
)

func ContactRouter(router *gin.Engine){
	router.GET("/api/contacts", controller.GetAll)
	router.POST("/api/contacts", controller.Create)
	router.DELETE("/api/contacts", controller.DeleteAll)
	router.POST("/api/contacts/add", controller.Create10Row)
	
	router.GET("/api/contacts/:id", controller.GetOne)
	router.PUT("/api/contacts/:id", controller.Update)
	router.DELETE("/api/contacts/:id", controller.DeleteOne)
}