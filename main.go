package main

import (
	"context"
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Start.")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	usercollection := client.Database("instadupe").Collection("users")
	userController := createUserController(usercollection)

	postcollection := client.Database("instadupe").Collection("posts")
	postController := createPostController(postcollection)

	r := httprouter.New()

	// End points
	r.POST("/users", userController.addUser)
	r.GET("/users", userController.getUsers)
	r.GET("/users/:id", userController.getSingleUserbyID)
	r.DELETE("/users", userController.deleteUsers)

	r.POST("/posts", postController.addPost)
	r.GET("/posts", postController.getPosts)
	r.GET("/posts/:id", postController.getSinglePost)
	r.DELETE("/posts", postController.deleteposts)

	r.GET("/posts/:id/users", postController.getPostByUser)

	log.Fatal(http.ListenAndServe(":8080", r))

}
