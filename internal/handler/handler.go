package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"qa-service/internal/repository"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Handler struct {
	qr repository.QuestionRepository
	ar repository.AnswerRepository
}

func NewHandler(qr repository.QuestionRepository, ar repository.AnswerRepository) *Handler {
	return &Handler{qr: qr, ar: ar}
}

func parseID(s string) (uint, bool) {
	id, err := strconv.ParseUint(s, 10, 32)
	return uint(id), err == nil
}

func (h *Handler) respondJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func (h *Handler) respondError(w http.ResponseWriter, status int, msg string) {
	h.respondJSON(w, status, map[string]string{"error": msg})
}
