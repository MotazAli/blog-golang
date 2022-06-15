package routers

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func UsersRoute(router *gin.RouterGroup,DB *mongo.Client){
	usersController := controllers.CreateUsersController(DB)
	router.GET("",usersController.GetAllUsers())
	router.GET(":id",usersController.GetUserById())
	router.POST("", usersController.CreateUser())
	router.PUT(":id", usersController.UpdateUserById())
	router.DELETE(":id",usersController.DeleteUserById())
}