package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Posts struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AccountId        primitive.ObjectID `json:"accountid,omitempty" bson:"accountid,omitempty"`
	Caption          string             `json:"caption,omitempty" bson:"caption,omitempty"`
	Image_URL        string             `json:"imageurl,omitempty" bson:"imageurl,omitempty"`
	Posted_Timestamp time.Time          `json:"posted_Timestamp,omitempty" bson:"Posted_Timestamp,omitempty"`
}

type PostsController struct {
	postscollection *mongo.Collection
}

func createPostController(collection *mongo.Collection) *PostsController {
	return &PostsController{collection}
}
