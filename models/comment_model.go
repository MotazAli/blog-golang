package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Comment struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Body      string             `json:"body,omitempty" validate:"required"`
	User      UserLight 		 `json:"user,omitempty" validate:"required"`
	Post      PostLight 		 `json:"post_id,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CommentLight struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Body      string             `json:"body,omitempty" validate:"required"`
	User      UserLight 		 `json:"user,omitempty" validate:"required"`
}

type CommentCreateRequest struct {
	Body    string `json:"body,omitempty" validate:"required"`
	UserId  string `json:"user_id,omitempty" validate:"required"`
	PostId  string `json:"post_id,omitempty" validate:"required"`
}

type CommentUpdateRequest struct {
	Body    string `json:"body,omitempty" validate:"required"`
}

////////// for testing only ////////////////
type CommentTest struct {
	Id        string   		  `json:"id,omitempty"`
	Body      string   		  `json:"body,omitempty" validate:"required"`
	User      UserTest 		  `json:"user,omitempty" validate:"required"`
	Post	  PostMinimalTest `json:"post,omitempty" validate:"required"`
	CreatedAt string   		  `json:"created_at"`
	UpdatedAt string  		  `json:"updated_at"`
}