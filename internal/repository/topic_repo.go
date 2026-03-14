package repository

import (
	"database/sql"
	"librongo/internal/models"
)

type TopicRepository struct {
	db *sql.DB
}

func NewTopicRepository(db *sql.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

func (r *TopicRepository) GetAll() ([]models.Topic, error) {

	query := `
	SELECT id, name
	FROM topics
	ORDER BY name
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var topics []models.Topic

	for rows.Next() {
		var topic models.Topic

		err := rows.Scan(
			&topic.ID,
			&topic.Name,
		)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}

func (r *TopicRepository) Create(name string) (int64, error) {

	query := `
	INSERT INTO topics (name)
	VALUES (?)
	`

	result, err := r.db.Exec(query, name)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}
