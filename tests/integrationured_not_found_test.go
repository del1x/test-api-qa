package tests

import (
	"net/http"
	"net/http/httptest"
	"qa-service/internal/handler"
	"qa-service/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	db := setupTestDB()
	mux := http.NewServeMux()
	handler.SetupRoutes(mux, repository.NewQuestionRepository(db), repository.NewAnswerRepository(db))

	req := httptest.NewRequest(http.MethodGet, "/questions/999999", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
