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

type UsersRepository struct {
	DB *mongo.Client
}


func (repository UsersRepository) FindAllUsers() ([]models.UserLight,error){
	var userCollection *mongo.Collection = configs.GetCollection(repository.DB,"users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var users []models.UserLight = []models.UserLight{}
        defer cancel()

        results, err := userCollection.Find(ctx, bson.M{})

        if err != nil {
            return nil, err
        }
		defer results.Close(ctx)
        

		err = results.All(ctx,&users)
		if err != nil {
            return nil, err
        }

		return users,nil
}


func (repository UsersRepository) FindAllUsersPaging(page int, size int) ([]models.UserLight,error){
	var userCollection *mongo.Collection = configs.GetCollection(repository.DB,"users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var users []models.UserLight = []models.UserLight{}
        defer cancel()

		skip := int64(page)
		limit := int64(size)
		opts := options.FindOptions{
			Skip: &skip,
			Limit: &limit,
		  }
        results, err := userCollection.Find(ctx, bson.M{},&opts)

        if err != nil {
           return nil, err
        }
		defer results.Close(ctx)
        

		err = results.All(ctx,&users)
		if err != nil {
            return nil, err
        }
      

		return users,nil
}


func (repository UsersRepository) FindOneUserById(id string) (*models.User,error){
	var userCollection *mongo.Collection = configs.GetCollection(repository.DB,"users")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
	var user models.User
    err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
    if err != nil {
        return nil,err
    }

	return &user,nil
}

func (repository UsersRepository)InsertUser(newUser *models.User) (*models.User,error){
	var userCollection *mongo.Collection = configs.GetCollection(repository.DB,"users")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	
	defer cancelFunc()

	


	_, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return nil,err
	}

	return newUser,nil
}


func (repository UsersRepository) UpdateUser(id string, editUser *models.User)(*models.User,error){
	var userCollection *mongo.Collection = configs.GetCollection(repository.DB,"users")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)

	update := bson.M{"name": editUser.Name, "email": editUser.Email, "password": editUser.Password,"posts":editUser.Posts, "updated_at":editUser.UpdatedAt}
    _ , err1 := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
      
    if err1 != nil {
        return nil,err1
    }
	return editUser,nil

}
func (repository UsersRepository) DeleteUserById(id string) (*models.User,error){
	oldUser, err := repository.FindOneUserById(id)
	if err != nil{
		return nil,err
	}

	var userCollection *mongo.Collection = configs.GetCollection(repository.DB,"users")
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	objId, _ := primitive.ObjectIDFromHex(id)
    _ , err1 := userCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err1 != nil{
		return nil,err1
	}

	return oldUser,nil
}


