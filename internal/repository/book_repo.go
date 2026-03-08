package repository

import (
	"database/sql"
	"librongo/internal/models"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) BeginTx() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *BookRepository) GetByID(id int) (*models.Book, error) {
	query := `
	SELECT 
		b.id, b.name, b.year, b.language_code, b.isbn, b.edition,
		b.cover_url, b.pages, b.location,
		e.id, e.name,
		a.id, a.name,
		t.id, t.name
	FROM books b
	LEFT JOIN editorials e ON b.editorial_id = e.id
	LEFT JOIN book_authors ba ON b.id = ba.book_id
	LEFT JOIN authors a ON ba.author_id = a.id
	LEFT JOIN book_topics bt ON b.id = bt.book_id
	LEFT JOIN topics t ON bt.topic_id = t.id
	WHERE b.id = ?
	`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var book *models.Book
	authorMap := make(map[int]bool)
	topicMap := make(map[int]bool)

	for rows.Next() {
		var (
			bID, year, pages                                  int
			name, language, isbn, edition, coverURL, location string

			editorialID   sql.NullInt64
			editorialName sql.NullString

			authorID   sql.NullInt64
			authorName sql.NullString

			topicID   sql.NullInt64
			topicName sql.NullString
		)

		err := rows.Scan(
			&bID, &name, &year, &language, &isbn, &edition,
			&coverURL, &pages, &location,
			&editorialID, &editorialName,
			&authorID, &authorName,
			&topicID, &topicName,
		)
		if err != nil {
			return nil, err
		}

		if book == nil {
			book = &models.Book{
				ID:       bID,
				Name:     name,
				Year:     year,
				Language: language,
				ISBN:     isbn,
				Edition:  edition,
				CoverURL: coverURL,
				Pages:    pages,
				Location: location,
				Authors:  []models.Author{},
				Topics:   []models.Topic{},
			}

			if editorialID.Valid {
				book.Editorial = &models.Editorial{
					ID:   int(editorialID.Int64),
					Name: editorialName.String,
				}
			}
		}

		// Agregar autor sin duplicar
		if authorID.Valid && !authorMap[int(authorID.Int64)] {
			book.Authors = append(book.Authors, models.Author{
				ID:   int(authorID.Int64),
				Name: authorName.String,
			})
			authorMap[int(authorID.Int64)] = true
		}

		// Agregar topic sin duplicar
		if topicID.Valid && !topicMap[int(topicID.Int64)] {
			book.Topics = append(book.Topics, models.Topic{
				ID:   int(topicID.Int64),
				Name: topicName.String,
			})
			topicMap[int(topicID.Int64)] = true
		}
	}

	if book == nil {
		return nil, sql.ErrNoRows
	}

	return book, nil
}

func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	query := `
	SELECT 
		b.id, b.name, b.year, b.language_code, b.isbn, b.edition,
		b.cover_url, b.pages, b.location,
		e.id, e.name,
		a.id, a.name,
		t.id, t.name
	FROM books b
	LEFT JOIN editorials e ON b.editorial_id = e.id
	LEFT JOIN book_authors ba ON b.id = ba.book_id
	LEFT JOIN authors a ON ba.author_id = a.id
	LEFT JOIN book_topics bt ON b.id = bt.book_id
	LEFT JOIN topics t ON bt.topic_id = t.id
	ORDER BY b.id
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	booksMap := make(map[int]*models.Book)

	for rows.Next() {
		var (
			bID, year, pages                                  int
			name, language, isbn, edition, coverURL, location string

			editorialID   sql.NullInt64
			editorialName sql.NullString

			authorID   sql.NullInt64
			authorName sql.NullString

			topicID   sql.NullInt64
			topicName sql.NullString
		)

		err := rows.Scan(
			&bID, &name, &year, &language, &isbn, &edition,
			&coverURL, &pages, &location,
			&editorialID, &editorialName,
			&authorID, &authorName,
			&topicID, &topicName,
		)
		if err != nil {
			return nil, err
		}

		book, exists := booksMap[bID]
		if !exists {
			book = &models.Book{
				ID:       bID,
				Name:     name,
				Year:     year,
				Language: language,
				ISBN:     isbn,
				Edition:  edition,
				CoverURL: coverURL,
				Pages:    pages,
				Location: location,
				Authors:  []models.Author{},
				Topics:   []models.Topic{},
			}

			if editorialID.Valid {
				book.Editorial = &models.Editorial{
					ID:   int(editorialID.Int64),
					Name: editorialName.String,
				}
			}

			booksMap[bID] = book
		}

		if authorID.Valid {
			alreadyExists := false
			for _, a := range book.Authors {
				if a.ID == int(authorID.Int64) {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				book.Authors = append(book.Authors, models.Author{
					ID:   int(authorID.Int64),
					Name: authorName.String,
				})
			}
		}

		if topicID.Valid {
			alreadyExists := false
			for _, t := range book.Topics {
				if t.ID == int(topicID.Int64) {
					alreadyExists = true
					break
				}
			}
			if !alreadyExists {
				book.Topics = append(book.Topics, models.Topic{
					ID:   int(topicID.Int64),
					Name: topicName.String,
				})
			}
		}
	}

	// Convertir map a slice
	var books []models.Book
	for _, b := range booksMap {
		books = append(books, *b)
	}

	return books, nil
}

/* METODO PARA INSERTAR LIBRO BASE */
func (r *BookRepository) Create(tx *sql.Tx, book *models.BookCreate) (int64, error) {
	query := `
		INSERT INTO books 
		(name, year, language_code, isbn, edition, cover_url, editorial_id, pages, location)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	result, err := tx.Exec(
		query,
		book.Name,
		book.Year,
		book.Language,
		book.ISBN,
		book.Edition,
		book.CoverURL,
		book.EditorialID,
		book.Pages,
		book.Location,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// ======================
/* METODOS AUXILIARES */
// ======================

func (r *BookRepository) AddAuthors(tx *sql.Tx, bookID int64, authors []models.IDList) error {
	query := `INSERT INTO book_authors (book_id, author_id) VALUES (?, ?)`

	for _, author := range authors {
		_, err := tx.Exec(query, bookID, author.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *BookRepository) AddTopics(tx *sql.Tx, bookID int64, topics []models.IDList) error {
	query := `INSERT INTO book_topics (book_id, topic_id) VALUES (?, ?)`

	for _, topic := range topics {
		_, err := tx.Exec(query, bookID, topic.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
