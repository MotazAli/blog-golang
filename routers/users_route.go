package routers

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func UsersRoute(router *gin.RouterGroup,DB *mongo.Client){
	userController := controllers.CreateUsersController(DB)
	router.GET("",userController.GetAllUsers())
	router.GET(":id",userController.GetUserById())
	router.POST("", userController.CreateUser())
	router.PUT(":id", userController.UpdateUserById())
	router.DELETE(":id",userController.DeleteUserById())
}