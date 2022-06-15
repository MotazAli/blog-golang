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

	userId, _ := primitive.ObjectIDFromHex(newComment.UserId)
	postId, _ := primitive.ObjectIDFromHex(newComment.PostId)
    nowTime := time.Now() 
    comment := models.Comment{
        Id:primitive.NewObjectID(),
        Body: newComment.Body,
        UserId:userId,
		PostId: postId,
        CreatedAt: nowTime,
        UpdatedAt: nowTime,
    }
    return c.Repository.InsertComment(&comment)
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
    return  c.Repository.UpdateComment(id,comment)
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
	return comment,nil
}

