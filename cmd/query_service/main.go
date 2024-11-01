package main

import (
	"log"
	"net/http"
	"os"

	"context"
	"microblog/internal/query/handler"
	"microblog/internal/query/repository"

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
	//repo := repository.NewQueryRepository(db)
	repo := repository.NewQueryRepository(db)
	queryHandler := handler.NewQueryHandler(*repo)

	http.HandleFunc("/timeline", queryHandler.GetTimeline)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
