package service

import (
	"errors"
	"microblog/internal/query/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock del repositorio QueryRepositoryInterface
type MockQueryRepository struct {
	mock.Mock
}

func (m *MockQueryRepository) GetFollowedTweets(userID string) ([]model.Tweet, error) {
	args := m.Called(userID)
	// Usamos una comprobación condicional para retornar nil si no hay datos
	if tweets, ok := args.Get(0).([]model.Tweet); ok {
		return tweets, args.Error(1)
	}
	return nil, args.Error(1)
}

// TestQueryService_GetFollowedTweets verifica que el servicio retorne los tweets seguidos por un usuario.
func TestQueryService_GetFollowedTweets(t *testing.T) {
	mockRepo := new(MockQueryRepository)
	service := NewQueryService(mockRepo)

	userID := "user123"
	tweets := []model.Tweet{
		{ID: "1", UserID: "user456", Content: "Hello world!"},
		{ID: "2", UserID: "user789", Content: "Go is awesome!"},
	}

	mockRepo.On("GetFollowedTweets", userID).Return(tweets, nil)

	result, err := service.GetFollowedTweets(userID)

	assert.NoError(t, err)
	assert.Equal(t, tweets, result)

	mockRepo.AssertCalled(t, "GetFollowedTweets", userID)
}

// TestQueryService_GetFollowedTweets_Error verifica que el servicio retorne un error si el repositorio retorna un error.
func TestQueryService_GetFollowedTweets_Error(t *testing.T) {
	mockRepo := new(MockQueryRepository)
	service := NewQueryService(mockRepo)

	userID := "user123"
	mockRepo.On("GetFollowedTweets", userID).Return(nil, errors.New("database error"))

	result, err := service.GetFollowedTweets(userID)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "database error")

	mockRepo.AssertCalled(t, "GetFollowedTweets", userID)
}
