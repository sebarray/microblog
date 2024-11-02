package handler_test

import (
	"encoding/json"
	"errors"

	"microblog/internal/query/handler"
	"microblog/internal/query/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockQueryService simula el servicio QueryServiceInterface
type MockQueryService struct {
	mock.Mock
}

func (m *MockQueryService) GetFollowedTweets(userID string) ([]model.Tweet, error) {
	args := m.Called(userID)

	if tweets, ok := args.Get(0).([]model.Tweet); ok {
		return tweets, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestGetTimeline_Success(t *testing.T) {
	mockService := new(MockQueryService)
	handler := handler.NewQueryHandler(mockService)

	userID := "test_user"
	expectedTweets := []model.Tweet{
		{ID: "1", UserID: "test_user", Content: "Hello World!"},
		{ID: "2", UserID: "test_user", Content: "Another Tweet"},
	}

	// Configurar el mock para devolver tweets sin errores
	mockService.On("GetFollowedTweets", userID).Return(expectedTweets, nil)

	// Preparar la solicitud HTTP y el ResponseRecorder
	req, err := http.NewRequest("GET", "/timeline?userID="+userID, nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()

	// Llamar al handler
	handler.GetTimeline(rec, req)

	// Verificar el resultado
	assert.Equal(t, http.StatusOK, rec.Code)
	var resultTweets []model.Tweet
	err = json.Unmarshal(rec.Body.Bytes(), &resultTweets)
	assert.NoError(t, err)
	assert.Equal(t, expectedTweets, resultTweets)
	mockService.AssertCalled(t, "GetFollowedTweets", userID)
}

func TestGetTimeline_ServiceError(t *testing.T) {
	mockService := new(MockQueryService)
	handler := handler.NewQueryHandler(mockService)

	userID := "invalid_user"
	mockService.On("GetFollowedTweets", userID).Return(nil, errors.New("user not found"))

	req, err := http.NewRequest("GET", "/timeline?userID="+userID, nil)
	assert.NoError(t, err)
	rec := httptest.NewRecorder()

	handler.GetTimeline(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	assert.Contains(t, rec.Body.String(), "user not found")
	mockService.AssertCalled(t, "GetFollowedTweets", userID)
}
