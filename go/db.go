package swagger

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

func addActorToDB(db *sql.DB, actor Actor) (string, error) {
	// Проверяем, что все необходимые поля заполнены.
	if actor.Name == nil || actor.Gender == nil || actor.BirthDate == "" {
		return "", errors.New("не все поля актера заполнены")
	}

	// Подготовка SQL-запроса для вставки актера.
	query := `INSERT INTO actors (id, name, gender, birth_date, movies) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string
	err := db.QueryRow(query, actor.Id, actor.Name.FirstName, actor.Gender, actor.BirthDate, pq.Array(actor.Movies)).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("ошибка при добавлении актера в базу данных: %w", err)
	}

	return id, nil
}

func getActorFromDB(db *sql.DB, actorID string) (*Actor, error) {
	query := `SELECT id, name, gender, movies, birthDate FROM actors WHERE id = ?`
	row := db.QueryRow(query, actorID)

	var actor Actor
	err := row.Scan(&actor.Id, &actor.Name, &actor.Gender, pq.Array(&actor.Movies), &actor.BirthDate)
	if err != nil {
		return nil, err
	}

	return &actor, nil
}
