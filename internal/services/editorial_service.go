package services

import (
	"errors"
	"librongo/internal/models"
	"librongo/internal/repository"
)

type EditorialService struct {
	repo *repository.EditorialRepository
}

func NewEditorialService(repo *repository.EditorialRepository) *EditorialService {
	return &EditorialService{repo: repo}
}

func (s *EditorialService) GetAllEditorials() ([]models.Editorial, error) {
	return s.repo.GetAll()
}

func (s *EditorialService) CreateEditorial(name string) (int64, error) {

	if name == "" {
		return 0, errors.New("editorial name cannot be empty")
	}

	return s.repo.Create(name)
}
