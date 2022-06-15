package interfaces

import (
	"blog/models"

	//"github.com/gin-gonic/gin"
)

type IPostsRepository interface {
	FindAllPosts() ([]models.Post,error)
	FindAllPostsPaging(page int, size int) ([]models.Post,error)
	FindOnePostById(id string) (*models.Post,error)
	InsertPost(newPost *models.Post) (*models.Post,error)
	UpdatePost(id string, editPost *models.Post)(*models.Post,error)
	DeletePostById(id string) (int64,error)
}

type IPostsService interface {
	GetAllPosts() ([]models.Post,error)
	GetAllPostsPaging(page int, size int) ([]models.Post,error)
	GetPostById(id string) (*models.Post,error)
	CreatePost(newPost *models.PostRequest) (*models.Post,error)
	EditPost(id string, editPost *models.PostRequest)(*models.Post,error)
	RemovePostById(id string) (*models.Post,error)
}
