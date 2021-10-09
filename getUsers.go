package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

func (uc UserController) getUsers(response http.ResponseWriter, request *http.Request, p httprouter.Params) {
	response.Header().Add("content-type", "application/json")
	var UserArray []Users
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := uc.collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message: "` + err.Error() + `"}"`))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user Users
		cursor.Decode(&user)
		UserArray = append(UserArray, user)
	}

	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message: "` + err.Error() + `"}"`))
		return
	}
	json.NewEncoder(response).Encode(UserArray)
}
