package repository

import (
	"context"
	"microblog/internal/command/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func TestMain(m *testing.M) {
	// Configuración de la conexión a MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // Ajusta la URI según tu configuración
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	// Crea una base de datos para las pruebas
	db = client.Database("test_db")

	// Asegúrate de desconectar al finalizar las pruebas
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Ejecuta las pruebas
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateTweet(t *testing.T) {
	repo := NewCommandRepository(db)

	// Crea un tweet para la prueba
	tweet := &model.Tweet{
		Content: "Hello World",
	}

	// Llama al método CreateTweet
	err := repo.CreateTweet(tweet)

	// Verifica que no haya error
	assert.NoError(t, err)

	// Verifica que el tweet fue insertado en la base de datos
	count, err := db.Collection("tweets").CountDocuments(context.TODO(), map[string]interface{}{"content": tweet.Content})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count) // Asegúrate de que haya un documento
}

func TestFollowUser(t *testing.T) {
	repo := NewCommandRepository(db)

	// Llama al método FollowUser
	err := repo.FollowUser("followerID", "followeeID")

	// Verifica que no haya error
	assert.NoError(t, err)

	// Verifica que el seguimiento fue insertado en la base de datos
	count, err := db.Collection("followers").CountDocuments(context.TODO(), map[string]interface{}{"follower_id": "followerID", "followee_id": "followeeID"})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count) // Asegúrate de que haya un documento
}

func TestCleanUp(t *testing.T) {
	// Limpia las colecciones después de cada prueba
	_, err := db.Collection("tweets").DeleteMany(context.TODO(), map[string]interface{}{})
	assert.NoError(t, err)

	_, err = db.Collection("followers").DeleteMany(context.TODO(), map[string]interface{}{})
	assert.NoError(t, err)
}
