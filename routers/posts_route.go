package routers

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func PostsRoute(router *gin.RouterGroup,DB *mongo.Client){
	postsController := controllers.CreatePostsController(DB)
	router.GET("",postsController.GetAllPosts())
	router.GET(":id",postsController.GetPostById())
	router.POST("", postsController.CreatePost())
	router.PUT(":id", postsController.UpdatePostById())
	router.DELETE(":id",postsController.DeletePostById())
}