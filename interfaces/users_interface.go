package interfaces

import (
	"blog/models"

	//"github.com/gin-gonic/gin"
)

type IUsersRepository interface {
	FindAllUsers() ([]models.User,error)
	FindAllUsersPaging(page int, size int) ([]models.User,error)
	FindOneUserById(id string) (*models.User,error)
	InsertUser(newUser *models.User) (*models.User,error)
	UpdateUser(id string, editUser *models.User)(*models.User,error)
	DeleteUserById(id string) (*models.User,error)
}

type IUsersService interface {
	GetAllUsers() ([]models.User,error)
	GetAllUsersPaging(page int, size int) ([]models.User,error)
	GetUserById(id string) (*models.User,error)
	CreateUser(newUser *models.User) (*models.User,error)
	EditUser(id string, editUser *models.User)(*models.User,error)
	RemoveUserById(id string) (*models.User,error)
}
