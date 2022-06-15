package services

import (
	"blog/interfaces"
	"blog/models"
	"blog/utilities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type PostsService struct{
    Repository interfaces.IPostsRepository
}



func (p PostsService ) GetAllPosts()([]models.Post,error){	
        return p.Repository.FindAllPosts()
}

func (p PostsService ) GetAllPostsPaging(page int, size int) ([]models.Post,error){
    if page <= 0 { page = 1}
    page = size * (page - 1)
    return p.Repository.FindAllPostsPaging(page,size)
}

func (p PostsService) GetPostById(id string) (*models.Post,error){
    return p.Repository.FindOnePostById(id)
}

func (p PostsService ) CreatePost(newPost *models.PostRequest) (*models.Post,error){

    //validate required fields
    if validationErr := utilities.Validate.Struct(newPost); validationErr != nil {
        return nil,validationErr
    }

	userId, _ := primitive.ObjectIDFromHex(newPost.UserId)
    nowTime := time.Now() 
    post := models.Post{
        Id:primitive.NewObjectID(),
        Title:newPost.Title,
        Body: newPost.Body,
        UserId:userId,
        CreatedAt: nowTime,
        UpdatedAt: nowTime,
    }
    return p.Repository.InsertPost(&post)
}


func (p PostsService ) EditPost(id string, editPost *models.PostRequest)(*models.Post,error){
    
    //validate required fields
    if validationErr := utilities.Validate.Struct(editPost); validationErr != nil {
        return nil,validationErr
    }

    post, err := p.Repository.FindOnePostById(id)
    if err != nil{
        return nil,err
    }

    post.Title = editPost.Title
    post.Body = editPost.Body
    post.UpdatedAt = time.Now()
    return  p.Repository.UpdatePost(id,post)
}

func (p PostsService ) RemovePostById(id string) (*models.Post,error){
	post, err := p.Repository.FindOnePostById(id)
    if err != nil{
        return nil,err
    }
    _, err = p.Repository.DeletePostById(id)
	if err != nil{
        return nil,err
    }
	return post,nil
}

