package handler

import (
	"net/http"
	"strings"

	"qa-service/internal/repository"
)

func SetupRoutes(mux *http.ServeMux, qr repository.QuestionRepository, ar repository.AnswerRepository) {
	h := NewHandler(qr, ar)

	mux.HandleFunc("/questions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "/questions" || r.URL.Path == "/questions/" {
				h.ListQuestions(w, r)
				return
			}
			h.GetQuestion(w, r)
		case http.MethodPost:
			if r.URL.Path == "/questions" || r.URL.Path == "/questions/" {
				h.CreateQuestion(w, r)
				return
			}
			if strings.HasSuffix(r.URL.Path, "/answers") || strings.HasSuffix(r.URL.Path, "/answers/") {
				h.CreateAnswer(w, r)
				return
			}
			http.NotFound(w, r)
		case http.MethodDelete:
			h.DeleteQuestion(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/answers/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetAnswer(w, r)
		case http.MethodDelete:
			h.DeleteAnswer(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("QA Service API — ready!\n"))
	})
}
