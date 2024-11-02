package repository

import (
	"context"
	"log"

	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupTestDB(t *testing.T) *mongo.Database {
	// Conexión a MongoDB (puede estar en Docker)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Base de datos temporal para pruebas
	db := client.Database("microblog_test")

	// Limpieza de datos después del test
	t.Cleanup(func() {
		err := db.Drop(context.TODO())
		if err != nil {
			log.Printf("Error dropping test database: %v", err)
		}
		client.Disconnect(context.TODO())
	})

	return db
}

func seedTestData(db *mongo.Database, userID string) {
	// Inserta datos en la colección de seguidores
	followersCollection := db.Collection("followers")
	followersCollection.InsertMany(context.TODO(), []interface{}{
		bson.M{"follower_id": userID, "followed_id": "user1"},
		bson.M{"follower_id": userID, "followed_id": "user2"},
	})

	// Inserta datos en la colección de tweets
	tweetsCollection := db.Collection("tweets")
	tweetsCollection.InsertMany(context.TODO(), []interface{}{
		bson.M{"_id": "1", "user_id": "user1", "content": "Tweet from user1"},
		bson.M{"_id": "2", "user_id": "user2", "content": "Tweet from user2"},
	})
}

func TestGetFollowedTweets_Integration(t *testing.T) {
	db := setupTestDB(t)
	seedTestData(db, "test_user")

	// Crear el repositorio
	repo := NewQueryRepository(db)

	// Ejecutar el método a probar
	tweets, err := repo.GetFollowedTweets("test_user")

	// Verificar que no haya errores
	assert.NoError(t, err)

	// Verificar que se hayan obtenido los tweets correctos
	assert.Len(t, tweets, 2)
	assert.Equal(t, "Tweet from user1", tweets[0].Content)
	assert.Equal(t, "Tweet from user2", tweets[1].Content)
}
