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

// Fetch post based on the given post-id
func (pc PostsController) getSinglePost(response http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	response.Header().Add("content-type", "application/json")
	id_string := parameters.ByName("id")

	id, _ := primitive.ObjectIDFromHex(id_string)
	var post Posts
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := pc.postscollection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message: "` + err.Error() + `"}"`))
		return
	}
	json.NewEncoder(response).Encode(post)
}
