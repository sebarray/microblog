package handler

import (
	"errors"
	"microblog/internal/command/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCommandService es un mock de la interfaz CommandServiceInterface
type MockCommandService struct {
	mock.Mock
}

// Implementa el método CreateTweet del mock
func (m *MockCommandService) CreateTweet(tweet *model.Tweet) error {
	args := m.Called(tweet)
	return args.Error(0)
}

// Implementa el método FollowUser del mock
func (m *MockCommandService) FollowUser(followerID, followeeID string) error {
	args := m.Called(followerID, followeeID)
	return args.Error(0)
}

func TestPostTweet_Success(t *testing.T) {
	mockService := new(MockCommandService)
	handler := NewCommandHandler(mockService)

	// Configura el mock para el caso exitoso
	tweet := &model.Tweet{Content: "Hello World"}
	mockService.On("CreateTweet", tweet).Return(nil)

	// Crea una solicitud HTTP
	body := `{"content":"Hello World"}`
	req := httptest.NewRequest(http.MethodPost, "/tweets", strings.NewReader(body))
	w := httptest.NewRecorder()

	// Llama al método PostTweet
	handler.PostTweet(w, req)

	// Verifica la respuesta
	res := w.Result()
	assert.Equal(t, http.StatusCreated, res.StatusCode) // Verifica el código de estado
	mockService.AssertExpectations(t)                   // Verifica que se llamaron las expectativas
}

func TestPostTweet_Error(t *testing.T) {
	mockService := new(MockCommandService)
	handler := NewCommandHandler(mockService)

	// Configura el mock para devolver un error al crear el tweet
	tweet := &model.Tweet{Content: "Hello World"}
	mockService.On("CreateTweet", tweet).Return(errors.New("error creating tweet"))

	// Crea una solicitud HTTP
	body := `{"content":"Hello World"}`
	req := httptest.NewRequest(http.MethodPost, "/tweets", strings.NewReader(body))
	w := httptest.NewRecorder()

	// Llama al método PostTweet
	handler.PostTweet(w, req)

	// Verifica la respuesta
	res := w.Result()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode) // Verifica el código de estado
	mockService.AssertExpectations(t)                               // Verifica que se llamaron las expectativas
}
