package interfaces

import (
	"blog/models"
)

type ICommentsRepository interface {
	FindAllComments() ([]models.Comment,error)
	FindAllCommentsPaging(page int, size int) ([]models.Comment,error)
	FindOneCommentById(id string) (*models.Comment,error)
	InsertComment(newComment *models.Comment) (*models.Comment,error)
	UpdateComment(id string, editComment *models.Comment)(*models.Comment,error)
	DeleteCommentById(id string) (int64,error)
}

type ICommentsService interface {
	GetAllComments() ([]models.Comment,error)
	GetAllCommentsPaging(page int, size int) ([]models.Comment,error)
	GetCommentById(id string) (*models.Comment,error)
	CreateComment(newComment *models.CommentCreateRequest) (*models.Comment,error)
	EditComment(id string, editComment *models.CommentUpdateRequest)(*models.Comment,error)
	RemoveCommentById(id string) (*models.Comment,error)
}
