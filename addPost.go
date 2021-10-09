package main

import (
	"context"
	"net/http"
	"time"

	"encoding/json"

	"github.com/julienschmidt/httprouter"
)

// Add post information to the post database
func (pc PostsController) addPost(response http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	response.Header().Add("content-type", "application/json")
	var post Posts
	json.NewDecoder(request.Body).Decode(&post)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := pc.postscollection.InsertOne(ctx, post)
	json.NewEncoder(response).Encode(result)
}
