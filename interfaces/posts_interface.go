package interfaces

import (
	"blog/models"

	//"github.com/gin-gonic/gin"
)

type IPostsRepository interface {
	FindAllPosts() ([]models.PostLight,error)
	FindAllPostsPaging(page int, size int) ([]models.PostLight,error)
	FindOnePostById(id string) (*models.Post,error)
	InsertPost(newPost *models.Post) (*models.Post,error)
	UpdatePost(id string, editPost *models.Post)(*models.Post,error)
	DeletePostById(id string) (int64,error)
}

type IPostsService interface {
	GetAllPosts() ([]models.PostLight,error)
	GetAllPostsPaging(page int, size int) ([]models.PostLight,error)
	GetPostById(id string) (*models.Post,error)
	CreatePost(newPost *models.PostCreateRequest) (*models.Post,error)
	EditPost(id string, editPost *models.PostUpdateRequest)(*models.Post,error)
	RemovePostById(id string) (*models.Post,error)
	CreatePostCommentByPostId(id string,comment *models.CommentLight) (*models.Post,error)
	EditPostCommentByPostId(id string,comment *models.CommentLight) (*models.Post,error)
	DeletePostCommentByPostId(id string,commentId string) (*models.Post,error)
}
