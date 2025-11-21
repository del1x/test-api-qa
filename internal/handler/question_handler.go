package handler

import (
	"log/slog"
	"net/http"
)

func (h *Handler) ListQuestions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "ListQuestions — будет реализовано через 5 минут"}`))
	slog.Info("ListQuestions called")
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Question created"}`))
}

func (h *Handler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GetQuestion"}`))
}

func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// Answer handlers — тоже пока заглушки
func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Answer created"}`))
}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GetAnswer"}`))
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
