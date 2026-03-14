package handler

import (
	"database/sql"
	"encoding/json"
	"librongo/internal/models"
	"librongo/internal/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AuthorHandler struct {
	service *services.AuthorService
}

func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{
		service: service,
	}
}

func (h *AuthorHandler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {

	authors, err := h.service.GetAllAuthors()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (h *AuthorHandler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {

	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	author, err := h.service.GetAuthorByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "No hay autor con ese id", http.StatusBadRequest)
			return
		}
		http.Error(w, "ojo", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func (h *AuthorHandler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateAuthor(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"id": id})
}
