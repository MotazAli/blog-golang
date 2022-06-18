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


type CommentsController struct{
    Service interfaces.ICommentsService
}




// CreateComment godoc
// @Summary      create new comment 
// @Description  create new comment 
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        comment  body     models.CommentCreateRequest  true  "Add comment"
// @Success      201  {object}  responses.Response{data=models.Comment}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /comments [post]
func (controller CommentsController) CreateComment() gin.HandlerFunc{
	return func(c *gin.Context) {
		var comment models.CommentCreateRequest
        if err := c.BindJSON(&comment); err != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }
	
		result, err := controller.Service.CreateComment(&comment)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Result: map[string]interface{}{"data": result}})
	}
}


// GetAllComments godoc
// @Summary      Get all comments or get comments using pagination
// @Description  Get all comments or get comments using pagination
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        page   query      int  false  "Page number"
// @Param        size   query      int  false  "Number of object you want to return"
// @Success      200  {object}  responses.Response{data=[]models.Comment}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response 
// @Router       /comments [get]
func (controller CommentsController) GetAllComments() gin.HandlerFunc {
    return func(c *gin.Context){

        var result []models.Comment
        var err error
        size := c.Query("size") 
        page := c.Query("page")

        if size != "" && page != ""{
            sizeInt, _ := strconv.Atoi(size)
            pageInt, _ := strconv.Atoi(page)
            result, err = controller.Service.GetAllCommentsPaging(pageInt,sizeInt)

        } else {
            result, err = controller.Service.GetAllComments()
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

// GetCommentById   godoc
// @Summary      Get comment info by id
// @Description  Get comment info by id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Comment ID"
// @Success      200  {object}  responses.Response{data=models.Comment}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /comments/{id} [get]
func (controller CommentsController) GetCommentById() gin.HandlerFunc {
    return func(c *gin.Context){

        commentId := c.Param("id")
        result, err := controller.Service.GetCommentById(commentId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}},
        )

    }
}

// DeleteCommentById godoc
// @Summary      Delete comment by id 
// @Description  Delete comment by id
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Comment ID"
// @Success      200  {object}  responses.Response{data=models.Comment}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /comments/{id} [delete]
func (controller CommentsController) DeleteCommentById() gin.HandlerFunc{
	return func(c *gin.Context) {
		commentId := c.Param("id")
        result, err := controller.Service.RemoveCommentById(commentId)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK,
            responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}},
        )
	}
}

// UpdateCommentById godoc
// @Summary      Update comment by id 
// @Description  Update comment by id
// @Tags         comments
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Comment ID"
// @Param        comment  body      models.CommentUpdateRequest  true  "Update post"
// @Success      200  {object}  responses.Response{data=models.Comment}
// @Failure      400  {object}  responses.Response
// @Failure      500  {object}  responses.Response
// @Router       /comments/{id} [put]
func (controller CommentsController) UpdateCommentById() gin.HandlerFunc{
	return func(c *gin.Context) {
        commentId := c.Param("id")
		var comment models.CommentUpdateRequest
		
        if err := c.BindJSON(&comment); err != nil {
            c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

		result, err := controller.Service.EditComment(commentId,&comment)
        if err != nil {
            c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Result: map[string]interface{}{"data": err.Error()}})
            return
        }

        c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Result: map[string]interface{}{"data": result}})
	}
}


var commentsService *services.CommentsService = nil
func GetCommentsService(DB *mongo.Client) *services.CommentsService{
    if commentsService == nil{
        commentsRepository := repositories.CommentsRepository{DB:DB}
        postsService := GetPostsService(DB)
        usersService := GetUsersService(DB)
        commentsService = &services.CommentsService{Repository:commentsRepository,PostsService:postsService,UsersService: usersService }
    }
    return commentsService 
}


func CreateCommentsController(DB *mongo.Client) *CommentsController{
    commentsServiceObj := GetCommentsService(DB)
    return &CommentsController{Service:commentsServiceObj} 
} 

