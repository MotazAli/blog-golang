package services

import (
	"blog/interfaces"
	"blog/models"
	"blog/utilities"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type PostsService struct{
    Repository interfaces.IPostsRepository
    UsersService interfaces.IUsersService
}



func (p PostsService ) GetAllPosts()([]models.PostLight,error){	
        return p.Repository.FindAllPosts()
}

func (p PostsService ) GetAllPostsPaging(page int, size int) ([]models.PostLight,error){
    if page <= 0 { page = 1}
    page = size * (page - 1)
    return p.Repository.FindAllPostsPaging(page,size)
}

func (p PostsService) GetPostById(id string) (*models.Post,error){
    return p.Repository.FindOnePostById(id)
}



func (p PostsService ) CreatePost(newPost *models.PostCreateRequest) (*models.Post,error){

    //validate required fields
    if validationErr := utilities.Validate.Struct(newPost); validationErr != nil {
        return nil,validationErr
    }

	
    user, err := p.UsersService.GetUserById(newPost.UserId)
    if err != nil{
        return nil,err
    }

    userlight := models.UserLight{
        Id: user.Id,
        Name: user.Name,
        Email: user.Email,
    }

    nowTime := time.Now() 
    post := models.Post{
        Id:primitive.NewObjectID(),
        Title:newPost.Title,
        Body: newPost.Body,
        User: userlight,
        CreatedAt: nowTime,
        UpdatedAt: nowTime,
    }
    _ , err = p.Repository.InsertPost(&post)
    if err != nil{
        return nil,err
    }

    postMinimal := models.PostMinimal{
        Id: post.Id,
        Body: post.Body,
        Title: post.Title,
        CreatedAt:post.CreatedAt,
        UpdatedAt: post.UpdatedAt,
    }

    _, err = p.UsersService.CreateUserPostByUserId(newPost.UserId,&postMinimal)
    if err != nil{
        return nil,err
    }

    return &post,nil
}


func (p PostsService ) EditPost(id string, editPost *models.PostUpdateRequest)(*models.Post,error){
    
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
    _, err = p.Repository.UpdatePost(id,post)
    if err != nil{
        return nil,err
    }

    postMinimal := models.PostMinimal{
        Id: post.Id,
        Body: post.Body,
        Title: post.Title,
        CreatedAt:post.CreatedAt,
        UpdatedAt: post.UpdatedAt,
    }

    _, err = p.UsersService.EditUserPostByUserId(post.User.Id.Hex(),&postMinimal)
    if err != nil{
        return nil,err
    }

    return post,nil
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


    _, err = p.UsersService.DeleteUserPostByUserId(post.User.Id.Hex(),id)
    if err != nil{
        return nil,err
    }
	return post,nil
}

func (p PostsService ) CreatePostCommentByPostId(id string,comment *models.CommentLight) (*models.Post,error){
    post, err := p.Repository.FindOnePostById(id)
    if err != nil{
        return nil,err
    }

    post.Comments = append(post.Comments, *comment)
    return p.Repository.UpdatePost(id,post)
}

func (p PostsService ) EditPostCommentByPostId(id string,comment *models.CommentLight) (*models.Post,error){
    post, err := p.Repository.FindOnePostById(id)
    if err != nil{
        return nil,err
    }

    if post.Comments == nil{
        return nil,errors.New("No comments in post id " + id)
    }

    for i := 0; i < len(post.Comments); i++ {
        if post.Comments[i].Id == comment.Id{
            post.Comments[i].Body = comment.Body
            break
        }
    }
    
    return p.Repository.UpdatePost(id,post)
}


func (p PostsService ) DeletePostCommentByPostId(id string,commentId string) (*models.Post,error){
    post, err := p.Repository.FindOnePostById(id)
    if err != nil{
        return nil,err
    }

    if post.Comments == nil{
        return nil,errors.New("No comments in post id " + id)
    }

    for i := 0; i < len(post.Comments); i++ {
        if post.Comments[i].Id.Hex() == commentId{
            // post.Comments = removeCommentElementByIndex(post.Comments,i)
            post.Comments = utilities.GRemoveElementByIndex(post.Comments,i)
            break
        }
    }
    
    return p.Repository.UpdatePost(id,post)
}

// func removeCommentElementByIndex(s []models.CommentLight, index int) []models.CommentLight {
//     return append(s[:index], s[index+1:]...)
// }

// func GRemoveElementByIndex[E []E](s E, index int) E {
//     return append(s[:index], s[index+1:]...)
// }

