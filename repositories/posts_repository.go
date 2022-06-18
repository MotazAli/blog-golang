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

var postsCollection *mongo.Collection = nil
func (repository PostsRepository) getCollection() *mongo.Collection {
	if postsCollection == nil{
		postsCollection = configs.GetCollection(repository.DB,"posts")
	} 
	return postsCollection
}


func (repository PostsRepository) FindAllPosts() ([]models.PostLight,error){
	collection := repository.getCollection() 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var posts []models.PostLight = []models.PostLight{}
    defer cancel()

    results, err := collection.Find(ctx, bson.M{})
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


func (repository PostsRepository) FindAllPostsPaging(page int, size int) ([]models.PostLight,error){
	collection := repository.getCollection() 
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var posts []models.PostLight = []models.PostLight{}
    defer cancel()

	skip := int64(page)
	limit := int64(size)
	opts := options.FindOptions{
		Skip: &skip,
		Limit: &limit,
	  }
    results, err := collection.Find(ctx, bson.M{},&opts)

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
	collection := repository.getCollection() 
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
	var post models.Post
    err := collection.FindOne(ctx, bson.M{"id": objId}).Decode(&post)
    if err != nil {
        return nil,err
    }

	return &post,nil
}

func (repository PostsRepository)InsertPost(newPost *models.Post) (*models.Post,error){
	collection := repository.getCollection() 
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)	
	defer cancelFunc()

	_, err := collection.InsertOne(ctx, newPost)
	if err != nil {
		return nil,err
	}

	return newPost,nil
}


func (repository PostsRepository) UpdatePost(id string, editPost *models.Post)(*models.Post,error){
	collection := repository.getCollection() 
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"title": editPost.Title, "body": editPost.Body,"comments":editPost.Comments , "updated_at":editPost.UpdatedAt}
    _ , err1 := collection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
      
    if err1 != nil {
        return nil,err1
    }
	return editPost,nil

}
func (repository PostsRepository) DeletePostById(id string) (int64,error){
	collection := repository.getCollection() 
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
    result , err := collection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil{
		return -1,err
	}

	return result.DeletedCount,nil
}


