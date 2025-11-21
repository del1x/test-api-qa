package handler

import (
	"net/http"

	"qa-service/internal/repository"
)

type Handler struct {
	questionRepo repository.QuestionRepository
	answerRepo   repository.AnswerRepository
}

func NewHandler(qr repository.QuestionRepository, ar repository.AnswerRepository) *Handler {
	return &Handler{
		questionRepo: qr,
		answerRepo:   ar,
	}
}

func SetupRoutes(mux *http.ServeMux, questionRepo repository.QuestionRepository, answerRepo repository.AnswerRepository) {
	h := NewHandler(questionRepo, answerRepo)

	mux.HandleFunc("GET /questions/", h.ListQuestions)
	mux.HandleFunc("POST /questions/", h.CreateQuestion)
	mux.HandleFunc("GET /questions/{id}", h.GetQuestion)
	mux.HandleFunc("DELETE /questions/{id}", h.DeleteQuestion)

	mux.HandleFunc("POST /questions/{id}/answers/", h.CreateAnswer)
	mux.HandleFunc("GET /answers/{id}", h.GetAnswer)
	mux.HandleFunc("DELETE /answers/{id}", h.DeleteAnswer)

	// заглушка на корень
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("QA Service API is running!"))
	})
}
