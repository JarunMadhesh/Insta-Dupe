package main

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"

	"encoding/json"
)

func (pc PostsController) getPosts(response http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	response.Header().Add("content-type", "application/json")
	var postArray []Posts
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := pc.postscollection.Find(ctx, bson.M{})
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
