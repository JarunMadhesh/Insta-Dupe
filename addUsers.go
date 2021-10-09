package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Add user information to the user database
func (uc UserController) addUser(response http.ResponseWriter, request *http.Request, p httprouter.Params) {
	response.Header().Add("content-type", "application/json")
	var user Users
	json.NewDecoder(request.Body).Decode(&user)

	// Encrypting the password
	password := []byte(user.Password)
	key := []byte("This is a key for the very secret password")
	password, _ = encrypt(key, password)
	user.Password = string(password)
	//
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := uc.collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}
