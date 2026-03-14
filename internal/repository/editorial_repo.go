package repository

import (
	"database/sql"
	"librongo/internal/models"
)

type EditorialRepository struct {
	db *sql.DB
}

func NewEditorialRepository(db *sql.DB) *EditorialRepository {
	return &EditorialRepository{db: db}
}

func (r *EditorialRepository) GetAll() ([]models.Editorial, error) {

	query := `
	SELECT id, name
	FROM editorials
	ORDER BY name
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var editorials []models.Editorial

	for rows.Next() {
		var editorial models.Editorial

		err := rows.Scan(
			&editorial.ID,
			&editorial.Name,
		)
		if err != nil {
			return nil, err
		}

		editorials = append(editorials, editorial)
	}

	return editorials, nil
}

func (r *EditorialRepository) Create(name string) (int64, error) {

	query := `
	INSERT INTO editorials (name)
	VALUES (?)
	`

	result, err := r.db.Exec(query, name)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
