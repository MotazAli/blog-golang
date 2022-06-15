package servers

import (
	"blog/configs"
	"blog/docs"
	"blog/routers"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func BlogServerRun(){
	docs.SwaggerInfo.Title = "Blog API"
	docs.SwaggerInfo.Description = "This is a blog server with mongodb."
	docs.SwaggerInfo.Version = "1.0"
	//docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = "localhost:8080"

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	DB := configs.ConnectDB()
	
	v1 := router.Group("/api/v1")
	{
		usersRouterGroup := v1.Group("/users")
		routers.UsersRoute(usersRouterGroup,DB) // users router

		postsRouterGroup := v1.Group("/posts")
		routers.PostsRoute(postsRouterGroup,DB) // posts router

		commentsRouterGroup := v1.Group("/comments")
		routers.CommentsRoute(commentsRouterGroup,DB) // comments router
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run(":8080")
}
