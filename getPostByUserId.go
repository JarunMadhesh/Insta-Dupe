package main

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"encoding/json"
)

// lists all the posts posted by a particular user
func (pc PostsController) getPostByUser(response http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	response.Header().Add("content-type", "application/json")
	id_string := parameters.ByName("id")

	id, _ := primitive.ObjectIDFromHex(id_string)
	var postArray []Posts
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := pc.postscollection.Find(ctx, bson.M{"accountid": id})

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message: "` + err.Error() + `"}"`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var post Posts
		cursor.Decode(&post)
		postArray = append(postArray, post)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message: "` + err.Error() + `"}"`))
		return
	}
	json.NewEncoder(response).Encode(postArray)
}
