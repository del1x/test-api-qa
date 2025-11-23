package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"qa-service/internal/model"
)

type CreateQuestionRequest struct {
	Text string `json:"text" validate:"required,min=1,max=1000"`
}

func (h *Handler) ListQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.qr.GetAll()
	if err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to get questions")
		return
	}
	h.respondJSON(w, http.StatusOK, questions)
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	var req CreateQuestionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid json")
		return
	}
	if err := validate.Struct(req); err != nil {
		h.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	q := &model.Question{Text: req.Text}
	if err := h.qr.Create(q); err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to create")
		return
	}

	h.respondJSON(w, http.StatusCreated, q)
}

func (h *Handler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/questions/")
	if idStr == "" || idStr == "/" {
		h.ListQuestions(w, r)
		return
	}
	id, ok := parseID(idStr)
	if !ok {
		h.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	q, err := h.qr.GetByIDWithAnswers(id)
	if err != nil {
		h.respondError(w, http.StatusNotFound, "question not found")
		return
	}
	h.respondJSON(w, http.StatusOK, q)
}

func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/questions/")
	id, ok := parseID(idStr)
	if !ok {
		h.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.qr.Delete(id); err != nil {
		h.respondError(w, http.StatusNotFound, "question not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
