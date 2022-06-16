package services

import (
	"blog/interfaces"
	"blog/models"
	"blog/utilities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type CommentsService struct{
    Repository interfaces.ICommentsRepository
    PostsService interfaces.IPostsService
    UsersService interfaces.IUsersService
}



func (c CommentsService ) GetAllComments()([]models.Comment,error){	
        return c.Repository.FindAllComments()
}

func (c CommentsService ) GetAllCommentsPaging(page int, size int) ([]models.Comment,error){
    if page <= 0 { page = 1}
    page = size * (page - 1)
    return c.Repository.FindAllCommentsPaging(page,size)
}

func (c CommentsService) GetCommentById(id string) (*models.Comment,error){
    return c.Repository.FindOneCommentById(id)
}

func (c CommentsService ) CreateComment(newComment *models.CommentCreateRequest) (*models.Comment,error){

    //validate required fields
    if validationErr := utilities.Validate.Struct(newComment); validationErr != nil {
        return nil,validationErr
    }


    user,err := c.UsersService.GetUserById(newComment.UserId)
    if err != nil {
        return nil,err
    }
    
    post,err := c.PostsService.GetPostById(newComment.PostId)
    if err != nil {
        return nil,err
    }

    userlight := models.UserLight{
        Id: user.Id,
        Name: user.Name,
        Email: user.Email,
    }

    postlight := models.PostLight{
        Id: post.Id,
        Title: post.Title,
        Body: post.Body,
        User:  models.UserLight{
            Id: post.User.Id,
            Name: post.User.Name,
            Email: post.User.Email,
        } ,
    }

    nowTime := time.Now() 
    comment := models.Comment{
        Id:primitive.NewObjectID(),
        Body: newComment.Body,
        User:userlight,
		Post: postlight,
        CreatedAt: nowTime,
        UpdatedAt: nowTime,
    }
    result,err := c.Repository.InsertComment(&comment)
    if err != nil {
        return nil,err
    }

    commentLight := models.CommentLight{
        Id: comment.Id,
        Body: comment.Body,
        User: comment.User,
    }
    _,err = c.PostsService.CreatePostCommentByPostId(newComment.PostId,&commentLight)
    if err != nil {
        return nil,err
    }
    return result,nil
}


func (c CommentsService ) EditComment(id string, editComment *models.CommentUpdateRequest)(*models.Comment,error){
    
    //validate required fields
    if validationErr := utilities.Validate.Struct(editComment); validationErr != nil {
        return nil,validationErr
    }

    comment, err := c.Repository.FindOneCommentById(id)
    if err != nil{
        return nil,err
    }

    
    comment.Body = editComment.Body
    comment.UpdatedAt = time.Now()

    result,err := c.Repository.UpdateComment(id,comment)
    if err != nil {
        return nil,err
    }

    commentLight := models.CommentLight{
        Id: comment.Id,
        Body: comment.Body,
        User: comment.User,
    }
    _,err = c.PostsService.EditPostCommentByPostId(string(comment.Post.Id.Hex()),&commentLight)
    if err != nil {
        return nil,err
    }
    return result,nil
}

func (c CommentsService ) RemoveCommentById(id string) (*models.Comment,error){
	comment, err := c.Repository.FindOneCommentById(id)
    if err != nil{
        return nil,err
    }
    _, err = c.Repository.DeleteCommentById(id)
	if err != nil{
        return nil,err
    }

    _,err = c.PostsService.DeletePostCommentByPostId(comment.Post.Id.Hex(),comment.Id.Hex())
    if err != nil {
        return nil,err
    }
    
	return comment,nil
}

