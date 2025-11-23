package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"qa-service/internal/handler"
	"qa-service/internal/repository"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerFlow(t *testing.T) {
	db := setupTestDB()
	mux := http.NewServeMux()
	handler.SetupRoutes(mux, repository.NewQuestionRepository(db), repository.NewAnswerRepository(db))

	qPayload := []byte(`{"text":"Random question"}`)
	req := httptest.NewRequest(http.MethodPost, "/questions/", bytes.NewReader(qPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	var q map[string]any
	json.Unmarshal(w.Body.Bytes(), &q)
	questionID := uint(q["id"].(float64))

	aPayload := []byte(`{"user_id":"user-123","text":"42"}`)
	req = httptest.NewRequest(http.MethodPost, "/questions/"+strconv.Itoa(int(questionID))+"/answers/", bytes.NewReader(aPayload))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var answer map[string]any
	json.Unmarshal(w.Body.Bytes(), &answer)
	answerID := uint(answer["id"].(float64))

	req = httptest.NewRequest(http.MethodGet, "/answers/"+strconv.Itoa(int(answerID)), nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "42")
}
