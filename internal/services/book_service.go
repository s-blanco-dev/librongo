package services

import (
	"database/sql"
	"errors"
	"librongo/internal/models"
	"librongo/internal/repository"
)

type BookService struct {
	repo *repository.BookRepository
}

func NewBookService(repo *repository.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	book, err := s.repo.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Book not found")
		}
		return nil, err
	}

	return book, nil
}
