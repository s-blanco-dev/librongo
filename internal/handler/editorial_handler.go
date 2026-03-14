package handler

import (
	"encoding/json"
	"net/http"

	"librongo/internal/services"
)

type EditorialHandler struct {
	service *services.EditorialService
}

func NewEditorialHandler(service *services.EditorialService) *EditorialHandler {
	return &EditorialHandler{service: service}
}

func (h *EditorialHandler) GetAllEditorials(w http.ResponseWriter, r *http.Request) {

	editorials, err := h.service.GetAllEditorials()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(editorials)
}

func (h *EditorialHandler) CreateEditorial(w http.ResponseWriter, r *http.Request) {

	var editorial struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&editorial)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateEditorial(editorial.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int64{
		"id": id,
	})
}
