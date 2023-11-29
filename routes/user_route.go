package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.GET("/users", controllers.GetAllUsers())
	router.POST("/users", controllers.CreateUser())
	router.GET("/users/:userId", controllers.GetAUser())
	router.PATCH("/users/:userId", controllers.EditAUser())
	router.DELETE("/users/:userId", controllers.DeleteAUser())
}
