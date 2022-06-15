package routers

import (
	"blog/controllers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func CommentsRoute(router *gin.RouterGroup,DB *mongo.Client){
	commentsController := controllers.CreateCommentsController(DB)
	router.GET("",commentsController.GetAllComments())
	router.GET(":id",commentsController.GetCommentById())
	router.POST("", commentsController.CreateComment())
	router.PUT(":id", commentsController.UpdateCommentById())
	router.DELETE(":id",commentsController.DeleteCommentById())
}