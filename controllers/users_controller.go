package controllers

import (
	"blog/interfaces"
	"blog/models"
	"blog/repositories"
	"blog/responses"
	"strconv"

	"blog/services"


	"net/http"


	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


type UsersController struct{
    Service interfaces.IUsersService
}



// CreateUser godoc
// @Summary      create new user 
// @Description  create new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body     models.User  true  "Add user"
// @Success      201  {object}  responses.Response{data=models.User}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /users [post]
func (controller UsersController) CreateUser() gin.HandlerFunc{
	return func(c *gin.Context) {
		var user models.User
        
		
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }
	

		result, err := controller.Service.CreateUser(&user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": result}})
	}
}


// GetAllUsers godoc
// @Summary      Get all users or get users using pagination
// @Description  Get all users or get users using pagination
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number"
// @Param        size   query      int  false  "Number of object you want to return"
// @Success      200  {object}  responses.Response{data=[]models.UserLight}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response 
// @Router       /users [get]
func (controller UsersController) GetAllUsers() gin.HandlerFunc {
    return func(c *gin.Context){

        var result []models.UserLight
        var err error
        size := c.Query("size") 
        page := c.Query("page")

        if size != "" && page != ""{
            sizeInt, _ := strconv.Atoi(size)
            pageInt, _ := strconv.Atoi(page)
            result, err = controller.Service.GetAllUsersPaging(pageInt,sizeInt)

        } else {
            result, err = controller.Service.GetAllUsers()
        }

        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }
        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}},
        )
          
    }
}

// GetUserById   godoc
// @Summary      Get user info by id
// @Description  Get user info by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  responses.Response{data=models.User}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /users/{id} [get]
func (controller UsersController) GetUserById() gin.HandlerFunc {
    return func(c *gin.Context){

        userId := c.Param("id")
        result, err := controller.Service.GetUserById(userId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}},
        )

    }
}

// DeleteUserById godoc
// @Summary      Delete user by id 
// @Description  Delete user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  responses.Response{data=models.User}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /users/{id} [delete]
func (controller UsersController) DeleteUserById() gin.HandlerFunc{
	return func(c *gin.Context) {
		userId := c.Param("id")
        result, err := controller.Service.RemoveUserById(userId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}},
        )
	}
}

// UpdateUserById godoc
// @Summary      Update user by id 
// @Description  Update user by id
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Param        user  body      models.User  true  "Update user"
// @Success      200  {object}  responses.Response{data=models.User}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /users/{id} [put]
func (controller UsersController) UpdateUserById() gin.HandlerFunc{
	return func(c *gin.Context) {
        userId := c.Param("id")
		var user models.User
		
        if err := c.BindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

		result, err := controller.Service.EditUser(userId,&user)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}})
	}
}

var usersService *services.UsersService = nil
func GetUsersService(DB *mongo.Client) *services.UsersService{
    if usersService == nil{
        usersRepository := repositories.UsersRepository{DB:DB}
        usersService = &services.UsersService{Repository:usersRepository}
    }
    return usersService 
}

func CreateUsersController(DB *mongo.Client) *UsersController{
    usersServiceObj := GetUsersService(DB)
    return &UsersController{Service:usersServiceObj} 
} 




// func GetAllUsers() gin.HandlerFunc {
//     return services.GetAllUsers
// }


// func GetAllUsers() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//         var users []models.User
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
//     }
// }
