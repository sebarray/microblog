package main

import (
	"log"
	"net/http"
	"os"

	"context"
	"microblog/internal/query/handler"
	"microblog/internal/query/repository"
	"microblog/internal/query/service"

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

	repo := repository.NewQueryRepository(db)
	Service := service.NewQueryService(repo)
	queryHandler := handler.NewQueryHandler(*Service)

	http.HandleFunc("/timeline", queryHandler.GetTimeline)

	log.Println("Query service running on port 8081")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
