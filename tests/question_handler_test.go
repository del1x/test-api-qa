package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"qa-service/internal/handler"
	"qa-service/internal/repository"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=user password=pass dbname=qa_db port=5432 sslmode=disable",
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Exec("TRUNCATE questions, answers RESTART IDENTITY CASCADE")
	return db
}

func TestQuestionFlow(t *testing.T) {
	db := setupTestDB()
	mux := http.NewServeMux()
	handler.SetupRoutes(mux, repository.NewQuestionRepository(db), repository.NewAnswerRepository(db))

	payload := []byte(`{"text":"Randrom question 2?"}`)
	req := httptest.NewRequest(http.MethodPost, "/questions/", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var createdQuestion map[string]any
	json.Unmarshal(w.Body.Bytes(), &createdQuestion)
	questionID := uint(createdQuestion["id"].(float64))

	req = httptest.NewRequest(http.MethodGet, "/questions/"+strconv.Itoa(int(questionID)), nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	var gotQuestion map[string]any
	json.Unmarshal(w.Body.Bytes(), &gotQuestion)
	assert.Equal(t, "Randrom question 2?", gotQuestion["text"])

	req = httptest.NewRequest(http.MethodDelete, "/questions/"+strconv.Itoa(int(questionID)), nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
