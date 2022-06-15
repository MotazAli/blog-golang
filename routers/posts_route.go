package routers

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func PostsRoute(router *gin.RouterGroup,DB *mongo.Client){
	postController := controllers.CreatePostsController(DB)
	router.GET("",postController.GetAllPosts())
	router.GET(":id",postController.GetPostById())
	router.POST("", postController.CreatePost())
	router.PUT(":id", postController.UpdatePostById())
	router.DELETE(":id",postController.DeletePostById())
}