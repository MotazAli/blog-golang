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


type PostsController struct{
    Service interfaces.IPostsService
}




// CreatePost godoc
// @Summary      create post 
// @Description  create post 
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body     models.PostRequest  true  "Add post"
// @Success      201  {object}  responses.Response{data=models.Post}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /posts [post]
func (controller PostsController) CreatePost() gin.HandlerFunc{
	return func(c *gin.Context) {
		var post models.PostRequest
        if err := c.BindJSON(&post); err != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }
	
		result, err := controller.Service.CreatePost(&post)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": result}})
	}
}


// GetAllPosts godoc
// @Summary      Get all posts or get posts using pagination
// @Description  Get all posts or get posts using pagination
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number"
// @Param        size   query      int  false  "Number of object you want to return"
// @Success      200  {object}  responses.Response{data=[]models.Post}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response 
// @Router       /posts [get]
func (controller PostsController) GetAllPosts() gin.HandlerFunc {
    return func(c *gin.Context){

        var result []models.Post
        var err error
        size := c.Query("size") 
        page := c.Query("page")

        if size != "" && page != ""{
            sizeInt, _ := strconv.Atoi(size)
            pageInt, _ := strconv.Atoi(page)
            result, err = controller.Service.GetAllPostsPaging(pageInt,sizeInt)

        } else {
            result, err = controller.Service.GetAllPosts()
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

// GetPostById   godoc
// @Summary      Get post info by id
// @Description  Get post info by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      200  {object}  responses.Response{data=models.Post}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /posts/{id} [get]
func (controller PostsController) GetPostById() gin.HandlerFunc {
    return func(c *gin.Context){

        postId := c.Param("id")
        result, err := controller.Service.GetPostById(postId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}},
        )

    }
}

// DeletePostById godoc
// @Summary      Delete post by id 
// @Description  Delete post by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Success      200  {object}  responses.Response{data=models.Post}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /posts/{id} [delete]
func (controller PostsController) DeletePostById() gin.HandlerFunc{
	return func(c *gin.Context) {
		postId := c.Param("id")
        result, err := controller.Service.RemovePostById(postId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}},
        )
	}
}

// UpdatePostById godoc
// @Summary      Update post by id 
// @Description  Update post by id
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Post ID"
// @Param        post  body      models.PostRequest  true  "Update post"
// @Success      200  {object}  responses.Response{data=models.Post}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /posts/{id} [put]
func (controller PostsController) UpdatePostById() gin.HandlerFunc{
	return func(c *gin.Context) {
        postId := c.Param("id")
		var post models.PostRequest
		
        if err := c.BindJSON(&post); err != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

		result, err := controller.Service.EditPost(postId,&post)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": result}})
	}
}



func CreatePostsController(DB *mongo.Client) *PostsController{

    postsRepository := repositories.PostsRepository{DB:DB}
    postsService := services.PostsService{Repository:postsRepository}
    return &PostsController{Service:postsService} 
} 

