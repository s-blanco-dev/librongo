package handler

import (
	"encoding/json"
	"net/http"

	"librongo/internal/services"
)

type TopicHandler struct {
	service *services.TopicService
}

func NewTopicHandler(service *services.TopicService) *TopicHandler {
	return &TopicHandler{service: service}
}

func (h *TopicHandler) GetAllTopics(w http.ResponseWriter, r *http.Request) {

	topics, err := h.service.GetAllTopics()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(topics)
}

func (h *TopicHandler) CreateTopic(w http.ResponseWriter, r *http.Request) {

	var topic struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&topic)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateTopic(topic.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int64{
		"id": id,
	})
}
