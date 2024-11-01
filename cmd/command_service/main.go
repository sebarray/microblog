package main

import (
	"log"
	"net/http"
	"os"

	"context"
	"microblog/internal/command/handler"
	"microblog/internal/command/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURI := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("microblog")
	repo := repository.NewCommandRepository(db)
	commandHandler := handler.NewCommandHandler(*repo)

	http.HandleFunc("/tweet", commandHandler.PostTweet)
	http.HandleFunc("/follow", commandHandler.FollowUser)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
