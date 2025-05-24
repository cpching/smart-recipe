package test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cpching/smart-recipe/backend/internal/auth"
	"github.com/cpching/smart-recipe/backend/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type mockService struct{}

func (m *mockService) Register(ctx context.Context, email, password string) (domain.User, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return domain.User{Email: email, PasswordHash: string(hash)}, nil
}

type errorMockService struct{}

func (m *errorMockService) Register(ctx context.Context, email, password string) (domain.User, error) {
	return domain.User{}, errors.New("email already exists")
}

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	s := &mockService{}
	v := auth.NewValidation()
	h := auth.NewHandler(s, v)

	r := gin.Default()
	r.POST("/register", h.Register)
	return r
}

func setupTestRequest(r *gin.Engine, body map[string]string, method string, path string) *http.Request {
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, path, bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	return req
}

func setupTestRequestRaw(r *gin.Engine, json string, method string, path string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBufferString(json))
	req.Header.Set("Content-Type", "application/json")

	return req
}

func TestRegisterHandler_InvalidJson(t *testing.T) {
	r := setupTestRouter()
	invalidJSON := `{ "email": "test@example.com", "password": abc123-}` // missing quotes around password
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid request body")
}

func TestRegisterHandler_MissingEmail(t *testing.T) {
	r := setupTestRouter()
	body := map[string]string{
		"password": "Downl-123",
	}
	req := setupTestRequest(r, body, http.MethodPost, "/register")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "email is required")
}

func TestRegisterHandler_MissingPassword(t *testing.T) {
	r := setupTestRouter()
	body := map[string]string{
		"email": "test@example.com",
	}
	req := setupTestRequest(r, body, http.MethodPost, "/register")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password is required")
}

func TestRegisterHandler_InvalidEmail(t *testing.T) {
	r := setupTestRouter()
	body := map[string]string{
		"email":    "testexample.com",
		"password": "Downl-123",
	}
	req := setupTestRequest(r, body, http.MethodPost, "/register")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid email format")
}

func TestRegisterHandler_WeakPassword(t *testing.T) {
	r := setupTestRouter()
	body := map[string]string{
		"email":    "test@example.com",
		"password": "123", // too weak
	}

	req := setupTestRequest(r, body, http.MethodPost, "/register")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "password too weak")
}

func TestRegisterHandler_EmailAlreadyExists(t *testing.T) {
	r := setupTestRouter()
	body := map[string]string{
		"email":    "test@example.com",
		"password": "abcABC123-",
	}

	req := setupTestRequest(r, body, http.MethodPost, "/register")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	req = setupTestRequest(r, body, http.MethodPost, "/register")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "email already exists")
}

func TestRegisterHandler_Success(t *testing.T) {
	r := setupTestRouter()

	body := map[string]string{
		"email":    "test@example.com",
		"password": "abcABC123-",
	}

	req := setupTestRequest(r, body, http.MethodPost, "/register")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
