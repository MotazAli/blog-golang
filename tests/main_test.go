package tests

import (
	"blog/configs"
	"blog/controllers"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)



var RouterEngine *gin.Engine
var DbTest *mongo.Client
func TestMain(m *testing.M){

	RouterEngine = gin.Default()
	DbTest =configs.ConnectDBTest()
	configureUsersRoute(RouterEngine,DbTest)
	configurePostsRoute(RouterEngine,DbTest)
	configureCommentsRoute(RouterEngine,DbTest)
	rc := m.Run()
	os.Exit(rc)
}


func configureUsersRoute(*gin.Engine,*mongo.Client){
	usersController := controllers.CreateUsersController(DbTest)
	RouterEngine.GET("/api/v1/users",usersController.GetAllUsers())
	RouterEngine.GET("/api/v1/users/:id",usersController.GetUserById())
	RouterEngine.POST("/api/v1/users",usersController.CreateUser())
	RouterEngine.PUT("/api/v1/users/:id",usersController.UpdateUserById())
	RouterEngine.DELETE("/api/v1/users/:id",usersController.DeleteUserById())
}


func configurePostsRoute(*gin.Engine,*mongo.Client){
	postsController := controllers.CreatePostsController(DbTest)
	RouterEngine.GET("/api/v1/posts",postsController.GetAllPosts())
	RouterEngine.GET("/api/v1/posts/:id",postsController.GetPostById())
	RouterEngine.POST("/api/v1/posts",postsController.CreatePost())
	RouterEngine.PUT("/api/v1/posts/:id",postsController.UpdatePostById())
	RouterEngine.DELETE("/api/v1/posts/:id",postsController.DeletePostById())
}

func configureCommentsRoute(*gin.Engine,*mongo.Client){
	commentsController := controllers.CreateCommentsController(DbTest)
	RouterEngine.GET("/api/v1/comments",commentsController.GetAllComments())
	RouterEngine.GET("/api/v1/comments/:id",commentsController.GetCommentById())
	RouterEngine.POST("/api/v1/comments",commentsController.CreateComment())
	RouterEngine.PUT("/api/v1/comments/:id",commentsController.UpdateCommentById())
	RouterEngine.DELETE("/api/v1/comments/:id",commentsController.DeleteCommentById())
}




