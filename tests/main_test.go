package tests

import (
	"blog/configs"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var router *gin.Engine
var dbTest *mongo.Client
func GetRouter()*gin.Engine{
	if router == nil {
		router = gin.Default()
	}
	return router
	
}

func GetDatabase() *mongo.Client{
	if dbTest == nil {
		dbTest =configs.ConnectDBTest()
	}
	return dbTest
}



