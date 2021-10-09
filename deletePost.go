package main

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Deletes all the posts
// used while production and debugging
func (pc PostsController) deleteposts(response http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	response.Header().Add("content-type", "application/json")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pc.postscollection.Drop(ctx)
}
