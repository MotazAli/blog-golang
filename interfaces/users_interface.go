package interfaces

import (
	"blog/models"
)

type IUsersRepository interface {
	FindAllUsers() ([]models.UserLight,error)
	FindAllUsersPaging(page int, size int) ([]models.UserLight,error)
	FindOneUserById(id string) (*models.User,error)
	InsertUser(newUser *models.User) (*models.User,error)
	UpdateUser(id string, editUser *models.User)(*models.User,error)
	DeleteUserById(id string) (*models.User,error)
	
}

type IUsersService interface {
	GetAllUsers() ([]models.UserLight,error)
	GetAllUsersPaging(page int, size int) ([]models.UserLight,error)
	GetUserById(id string) (*models.User,error)
	CreateUser(newUser *models.User) (*models.User,error)
	EditUser(id string, editUser *models.User)(*models.User,error)
	RemoveUserById(id string) (*models.User,error)
	CreateUserPostByUserId(id string,newPost *models.PostMinimal) (*models.User,error)
	EditUserPostByUserId(id string,editPost *models.PostMinimal) (*models.User,error)
	DeleteUserPostByUserId(id string,postId string) (*models.User,error)
}
