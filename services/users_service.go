package services

import (
	//"blog/configs"
	"blog/interfaces"
	"blog/models"
	"time"

	//"blog/responses"
	//"context"
	//"net/http"
	//"time"

	//"github.com/gin-gonic/gin"
	//"go.mongodb.org/mongo-driver/bson"
	//"github.com/go-playground/validator"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo"
)


type UsersService struct{
    Repository interfaces.IUsersRepository
}


//var userCollection *mongo.Collection = configs.GetCollection(configs.DB,"users")
//var validate = validator.New()

var validate = validator.New()

func (u UsersService ) GetAllUsers()([]models.User,error){	
        return u.Repository.FindAllUsers()
}

func (u UsersService ) GetAllUsersPaging(page int, size int) ([]models.User,error){
    if page <= 0 { page = 1}
    page = size * (page - 1)
    return u.Repository.FindAllUsersPaging(page,size)
}

func (u UsersService) GetUserById(id string) (*models.User,error){
    return u.Repository.FindOneUserById(id)
}

func (u UsersService ) CreateUser(newUser *models.User) (*models.User,error){

    //validate required fields
    if validationErr := validate.Struct(newUser); validationErr != nil {
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
    if validationErr := validate.Struct(editUser); validationErr != nil {
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

// func GetAllUsers(c *gin.Context){
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//         var users []models.User = []models.User{}
//         defer cancel()

//         results, err := userCollection.Find(ctx, bson.M{})

//         if err != nil {
//             c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
//             return
//         }
// 		defer results.Close(ctx)
//         //reading from the db in an optimal way
        
//         for results.Next(ctx) {
//             var singleUser models.User
//             if err = results.Decode(&singleUser); err != nil {
//                 c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
//             }
          
//             users = append(users, singleUser)
//         }

//         c.JSON(http.StatusOK,
//             responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": users}},
//         )
// }