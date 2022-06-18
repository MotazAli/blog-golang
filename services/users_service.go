package services

import (
	"blog/interfaces"
	"blog/models"
	"blog/utilities"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type UsersService struct{
    Repository interfaces.IUsersRepository
}




func (u UsersService ) GetAllUsers()([]models.UserLight,error){	
        return u.Repository.FindAllUsers()
}

func (u UsersService ) GetAllUsersPaging(page int, size int) ([]models.UserLight,error){
    if page <= 0 { page = 1}
    page = size * (page - 1)
    return u.Repository.FindAllUsersPaging(page,size)
}

func (u UsersService) GetUserById(id string) (*models.User,error){
    return u.Repository.FindOneUserById(id)
}

func (u UsersService ) CreateUser(newUser *models.User) (*models.User,error){

    //validate required fields
    if validationErr := utilities.Validate.Struct(newUser); validationErr != nil {
        return nil,validationErr
    }

    nowTime := time.Now() 
    user := models.User{
        Id:primitive.NewObjectID(),
        Name:newUser.Name,
        Email: newUser.Email,
        Password:newUser.Password,
        CreatedAt: nowTime,
        UpdatedAt: nowTime,
    }
    return u.Repository.InsertUser(&user)
}


func (u UsersService ) EditUser(id string, editUser *models.User)(*models.User,error){
    
    //validate required fields
    if validationErr := utilities.Validate.Struct(editUser); validationErr != nil {
        return nil,validationErr
    }

    user, err := u.Repository.FindOneUserById(id)
    if err != nil{
        return nil,err
    }

    user.Name = editUser.Name
    user.Email = editUser.Email
    user.Password = editUser.Password
    user.UpdatedAt = time.Now()
    return u.Repository.UpdateUser(id,user)
}
func (u UsersService ) RemoveUserById(id string) (*models.User,error){
    return u.Repository.DeleteUserById(id)
}


func (u UsersService ) CreateUserPostByUserId(id string,newPost *models.PostMinimal) (*models.User,error){
    user, err := u.Repository.FindOneUserById(id)
    if err != nil{
        return nil,err
    }

    user.Posts = append(user.Posts, *newPost)
    return u.Repository.UpdateUser(id,user)
}
func (u UsersService ) EditUserPostByUserId(id string,editPost *models.PostMinimal) (*models.User,error){
    user, err := u.Repository.FindOneUserById(id)
    if err != nil{
        return nil,err
    }

    if user.Posts == nil{
        return nil,errors.New("No post with post id " + id )
    }

    for i := 0; i < len(user.Posts); i++ {
        if user.Posts[i].Id.Hex() == editPost.Id.Hex(){
            user.Posts[i].Body = editPost.Body
            user.Posts[i].Title = editPost.Title
            user.Posts[i].UpdatedAt = editPost.UpdatedAt
            break
        }
    }

    return u.Repository.UpdateUser(id,user)
}
func (u UsersService ) DeleteUserPostByUserId(id string,postId string) (*models.User,error){
    user, err := u.Repository.FindOneUserById(id)
    if err != nil{
        return nil,err
    }

    if user.Posts == nil{
        return nil,errors.New("No post with post id " + id)
    }

    for index , post := range user.Posts {
        if post.Id.Hex() == postId{
            // user.Posts = removePostElementByIndex(user.Posts,index)
            user.Posts = utilities.GRemoveElementByIndex(user.Posts,index)
            break
        }
    }

    return u.Repository.UpdateUser(id,user)
}

// func removePostElementByIndex(s []models.PostMinimal, index int) []models.PostMinimal {
//     return append(s[:index], s[index+1:]...)
// }
