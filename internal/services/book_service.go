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

func (s *BookService) CreateBook(book *models.BookCreate) (int64, error) {
	if book.Name == "" {
		return 0, errors.New("plana cabeza tiene usted")
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

	bookID, err := s.repo.Create(tx, book)
	if err != nil {
		return 0, err
	}

	if len(book.Authors) > 0 {
		err = s.repo.AddAuthors(tx, bookID, book.Authors)
		if err != nil {
			return 0, err
		}
	}

	if len(book.Topics) > 0 {
		err = s.repo.AddTopics(tx, bookID, book.Topics)
		if err != nil {
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return bookID, nil
}
