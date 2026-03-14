package services

import (
	"database/sql"
	"errors"
	"librongo/internal/models"
	"librongo/internal/repository"
)

type AuthorService struct {
	repo *repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: repo}
}

func (s *AuthorService) GetAllAuthors() ([]models.Author, error) {
	return s.repo.GetAll()
}

func (s *AuthorService) GetAuthorByID(id int) (*models.Author, error) {
	if id <= 0 {
		return nil, errors.New("Estás choto")
	}

	author, err := s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("No hay autor con ese id")
		}
		return nil, err
	}

	return author, nil
}

func (s *AuthorService) CreateAuthor(author *models.Author) (int64, error) {
	if author.Name == "" {
		return 0, errors.New("Autor debe tener nombre")
	}

	tx, err := s.repo.BeginTx()
	if err != nil {
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	authorID, err := s.repo.Create(tx, author)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return authorID, nil
}
