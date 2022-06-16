package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Post struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Body      string             `json:"body,omitempty" validate:"required"`
	User      UserLight 		 `json:"user,omitempty"`
	Comments  []CommentLight	 `json:"comments"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type PostLight struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Body      string             `json:"body,omitempty" validate:"required"`
	User      UserLight 		 `json:"user,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type PostMinimal struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title,omitempty" validate:"required"`
	Body      string             `json:"body,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}


type PostCreateRequest struct {
	Title  string `json:"title,omitempty" validate:"required"`
	Body   string `json:"body,omitempty" validate:"required"`
	UserId string `json:"user_id,omitempty" validate:"required"`
}

type PostUpdateRequest struct {
	Title  string `json:"title,omitempty" validate:"required"`
	Body   string `json:"body,omitempty" validate:"required"`
}