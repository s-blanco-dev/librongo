package repository

import (
	"database/sql"
	"librongo/internal/models"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) BeginTx() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *AuthorRepository) GetAll() ([]models.Author, error) {

	query := `
	SELECT id, name
	FROM authors
	ORDER BY name
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author

	for rows.Next() {
		var author models.Author

		err := rows.Scan(
			&author.ID,
			&author.Name,
		)
		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (r *AuthorRepository) GetByID(id int) (*models.Author, error) {
	query := `
	SELECT id, name FROM authors WHERE id = ?;
	`
	row := r.db.QueryRow(query, id)

	var author models.Author
	err := row.Scan(&author.ID, &author.Name)

	if err != nil {
		return nil, err
	}
	return &author, nil
}

func (r *AuthorRepository) Create(tx *sql.Tx, author *models.Author) (int64, error) {
	query := `
	INSERT INTO authors(name) VALUES(?);
	`

	result, err := tx.Exec(query, author.Name)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
