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

// Fetch user information based on the given user-id
func (uc UserController) getSingleUserbyID(response http.ResponseWriter, request *http.Request, parameters httprouter.Params) {
	response.Header().Add("content-type", "application/json")
	id_string := parameters.ByName("id")

	id, _ := primitive.ObjectIDFromHex(id_string)
	var user Users
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := uc.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message: "` + err.Error() + `"}"`))
		return
	}
	// Decrypting the password
	password := []byte(user.Password)
	key := []byte("This is a key for the very secret password")
	password, _ = decrypt(key, password)
	user.Password = string(password)
	//
	json.NewEncoder(response).Encode(user)
}
