// service/command_service_test.go
package service

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"microblog/internal/command/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCommandRepository es una estructura que simula CommandRepositoryInterface
type MockCommandRepository struct {
	mock.Mock
}

// Implementación del método CreateTweet para MockCommandRepository
func (m *MockCommandRepository) CreateTweet(tweet *model.Tweet) error {
	args := m.Called(tweet)
	return args.Error(0)
}

// Implementación del método FollowUser para MockCommandRepository
func (m *MockCommandRepository) FollowUser(followerID, followeeID string) error {
	args := m.Called(followerID, followeeID)
	return args.Error(0)
}

func TestCreateTweet(t *testing.T) {
	mockRepo := new(MockCommandRepository)
	service := NewCommandService(mockRepo)

	// Test caso exitoso
	tweet := &model.Tweet{Content: "Hello, world!"}
	mockRepo.On("CreateTweet", tweet).Return(nil)

	err := service.CreateTweet(tweet)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

	// Test caso de contenido demasiado largo
	longTweet := &model.Tweet{Content: strings.Repeat("A", 281)}
	err = service.CreateTweet(longTweet)

	assert.EqualError(t, err, "tweet content is too long")
	mockRepo.AssertNotCalled(t, "CreateTweet", longTweet)
}

func TestFollowUser(t *testing.T) {
	mockRepo := new(MockCommandRepository)
	serviceCommand := NewCommandService(mockRepo)

	// Configura el mock para una llamada exitosa a FollowUser
	mockRepo.On("FollowUser", "followerID", "followeeID").Return(nil)

	err := serviceCommand.FollowUser("followerID", "followeeID")
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)

}

func TestFollowUserError(t *testing.T) {
	mockRepo := new(MockCommandRepository)
	serviceCommand := NewCommandService(mockRepo)

	// Configura el mock para devolver un error al intentar seguir a un usuario
	mockRepo.On("FollowUser", mock.Anything, mock.Anything).Return(errors.New("error following user"))

	// Llamada al método de servicio
	err := serviceCommand.FollowUser("followerID", "followeeID")

	// Imprime el error para depuración
	fmt.Println("Error recibido:", err)

	// Verifica que se reciba el error esperado
	assert.EqualError(t, err, "error following user")

	// Asegúrate de que el mock haya sido llamado correctamente
	mockRepo.AssertExpectations(t)
}
