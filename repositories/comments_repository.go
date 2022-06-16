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

type CommentsRepository struct {
	DB *mongo.Client
}


func (repository CommentsRepository) FindAllComments() ([]models.Comment,error){
	var commentCollection *mongo.Collection = configs.GetCollection(repository.DB,"comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var comments []models.Comment = []models.Comment{}
    defer cancel()

    results, err := commentCollection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
	defer results.Close(ctx)
    
	err = results.All(ctx,&comments)
	if err != nil {
        return nil, err
    }
    
	return comments,nil
}


func (repository CommentsRepository) FindAllCommentsPaging(page int, size int) ([]models.Comment,error){
	var commentCollection *mongo.Collection = configs.GetCollection(repository.DB,"comments")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    var comments []models.Comment = []models.Comment{}
    defer cancel()

	skip := int64(page)
	limit := int64(size)
	opts := options.FindOptions{
		Skip: &skip,
		Limit: &limit,
	  }
    results, err := commentCollection.Find(ctx, bson.M{},&opts)

    if err != nil {
       return nil, err
    }
	defer results.Close(ctx)
    
	err = results.All(ctx,&comments)
	if err != nil {
        return nil, err
    }
	return comments,nil
}


func (repository CommentsRepository) FindOneCommentById(id string) (*models.Comment,error){
	var commentCollection *mongo.Collection = configs.GetCollection(repository.DB,"comments")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
	var comment models.Comment
    err := commentCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&comment)
    if err != nil {
        return nil,err
    }

	return &comment,nil
}

func (repository CommentsRepository)InsertComment(newComment *models.Comment) (*models.Comment,error){
	var commentCollection *mongo.Collection = configs.GetCollection(repository.DB,"comments")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)	
	defer cancelFunc()

	_, err := commentCollection.InsertOne(ctx, newComment)
	if err != nil {
		return nil,err
	}

	return newComment,nil
}


func (repository CommentsRepository) UpdateComment(id string, editComment *models.Comment)(*models.Comment,error){
	var commentCollection *mongo.Collection = configs.GetCollection(repository.DB,"comments")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{"body": editComment.Body, "updated_at":editComment.UpdatedAt}
    _ , err1 := commentCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
      
    if err1 != nil {
        return nil,err1
    }
	return editComment,nil

}
func (repository CommentsRepository) DeleteCommentById(id string) (int64,error){
	
	var commentCollection *mongo.Collection = configs.GetCollection(repository.DB,"comments")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
    result , err := commentCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil{
		return -1,err
	}

	return result.DeletedCount,nil
}


