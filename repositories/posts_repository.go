package repositories

import (
	"blog/configs"
	"blog/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostsRepository struct {
	DB *mongo.Client
}


func (repository PostsRepository) FindAllPosts() ([]models.Post,error){
	var postCollection *mongo.Collection = configs.GetCollection(repository.DB,"posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var posts []models.Post = []models.Post{}
    defer cancel()

    results, err := postCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
	defer results.Close(ctx)
    
	err = results.All(ctx,&posts)
	if err != nil {
        return nil, err
    }
    
	return posts,nil
}


func (repository PostsRepository) FindAllPostsPaging(page int, size int) ([]models.Post,error){
	var postCollection *mongo.Collection = configs.GetCollection(repository.DB,"posts")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var posts []models.Post = []models.Post{}
    defer cancel()

	skip := int64(page)
	limit := int64(size)
	opts := options.FindOptions{
		Skip: &skip,
		Limit: &limit,
	  }
    results, err := postCollection.Find(ctx, bson.M{},&opts)

    if err != nil {
       return nil, err
    }
	defer results.Close(ctx)
    
	err = results.All(ctx,&posts)
	if err != nil {
        return nil, err
    }
	return posts,nil
}


func (repository PostsRepository) FindOnePostById(id string) (*models.Post,error){
	var postCollection *mongo.Collection = configs.GetCollection(repository.DB,"posts")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
	var post models.Post
    err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&post)
    if err != nil {
        return nil,err
    }

	return &post,nil
}

func (repository PostsRepository)InsertPost(newPost *models.Post) (*models.Post,error){
	var postCollection *mongo.Collection = configs.GetCollection(repository.DB,"posts")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)	
	defer cancelFunc()

	_, err := postCollection.InsertOne(ctx, newPost)
	if err != nil {
		return nil,err
	}

	return newPost,nil
}


func (repository PostsRepository) UpdatePost(id string, editPost *models.Post)(*models.Post,error){
	var postCollection *mongo.Collection = configs.GetCollection(repository.DB,"posts")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"title": editPost.Title, "body": editPost.Body, "updated_at":editPost.UpdatedAt}
    _ , err1 := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
      
    if err1 != nil {
        return nil,err1
    }
	return editPost,nil

}
func (repository PostsRepository) DeletePostById(id string) (int64,error){
	// oldUser, err := repository.FindOneUserById(id)
	// if err != nil{
	// 	return nil,err
	// }

	var postCollection *mongo.Collection = configs.GetCollection(repository.DB,"posts")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
    result , err := postCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil{
		return -1,err
	}

	return result.DeletedCount,nil
}


