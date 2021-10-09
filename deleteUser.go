package main

import (
	"context"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Deletes all the users
// used while production and debugging
func (uc UserController) deleteUsers(response http.ResponseWriter, request *http.Request, p httprouter.Params) {
	response.Header().Add("content-type", "application/json")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uc.collection.Drop(ctx)
}
