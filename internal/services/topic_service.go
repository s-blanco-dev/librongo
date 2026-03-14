package services

import (
	"errors"
	"librongo/internal/models"
	"librongo/internal/repository"
)

type TopicService struct {
	repo *repository.TopicRepository
}

func NewTopicService(repo *repository.TopicRepository) *TopicService {
	return &TopicService{repo: repo}
}

func (s *TopicService) GetAllTopics() ([]models.Topic, error) {
	return s.repo.GetAll()
}

func (s *TopicService) CreateTopic(name string) (int64, error) {

	if name == "" {
		return 0, errors.New("topic name cannot be empty")
	}

	return s.repo.Create(name)
}
