package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Body      string             `json:"body,omitempty" validate:"required"`
	UserId    primitive.ObjectID `json:"user_id,omitempty" validate:"required"`
	PostId    primitive.ObjectID `json:"post_id,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CommentRequest struct {
	Body    string `json:"body,omitempty" validate:"required"`
	UserId  string `json:"user_id,omitempty" validate:"required"`
	PostId  string `json:"post_id,omitempty" validate:"required"`
}