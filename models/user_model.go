package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id 			primitive.ObjectID 	`json:"id,omitempty"`
	Name 		string 				`json:"name,omitempty" validate:"required"`
	Email 		string  			`json:"email,omitempty" validate:"required"`
	Password 	string  			`json:"password,omitempty" validate:"required"`
	Posts 		[]PostMinimal			`json:"posts"`
	CreatedAt 	time.Time 			`json:"created_at"`
	UpdatedAt 	time.Time 			`json:"updated_at"`
}

type UserLight struct {
	Id 			primitive.ObjectID 	`json:"id,omitempty"`
	Name 		string 				`json:"name,omitempty" validate:"required"`
	Email 		string  			`json:"email,omitempty" validate:"required"`
	CreatedAt 	time.Time 			`json:"created_at"`
	UpdatedAt 	time.Time 			`json:"updated_at"`
}