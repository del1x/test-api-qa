package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"qa-service/internal/model"

	"github.com/google/uuid"
)

type CreateAnswerRequest struct {
	UserID string `json:"user_id" validate:"required"`
	Text   string `json:"text" validate:"required,min=1,max=2000"`
}

func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/questions/")
	idStr = strings.TrimSuffix(idStr, "/answers/")
	id, ok := parseID(idStr)
	if !ok {
		h.respondError(w, http.StatusBadRequest, "invalid question id")
		return
	}

	var req CreateAnswerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid json")
		return
	}
	if err := validate.Struct(req); err != nil {
		h.respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := h.qr.GetByID(id); err != nil {
		h.respondError(w, http.StatusNotFound, "question not found")
		return
	}

	a := &model.Answer{
		QuestionID: id,
		UserID:     uuid.New().String(),
		Text:       req.Text,
	}
	if err := h.ar.Create(a); err != nil {
		h.respondError(w, http.StatusInternalServerError, "failed to create answer")
		return
	}

	h.respondJSON(w, http.StatusCreated, a)
}

func (h *Handler) GetAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/answers/")
	id, ok := parseID(idStr)
	if !ok {
		h.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}
	a, err := h.ar.GetByID(id)
	if err != nil {
		h.respondError(w, http.StatusNotFound, "answer not found")
		return
	}
	h.respondJSON(w, http.StatusOK, a)
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/answers/")
	id, ok := parseID(idStr)
	if !ok {
		h.respondError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.ar.Delete(id); err != nil {
		h.respondError(w, http.StatusNotFound, "answer not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
