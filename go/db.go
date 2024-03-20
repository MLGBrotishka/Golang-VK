package swagger

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

func addActorToDB(db *sql.DB, actor Actor) (string, error) {
	if actor.Name == nil || actor.Gender == nil || actor.BirthDate == "" {
		return "", errors.New("не все поля актера заполнены")
	}

	query := `INSERT INTO actors (id, name, gender, birth_date, movies) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id string
	err := db.QueryRow(query, actor.Id, actor.Name.FirstName, actor.Gender, actor.BirthDate, pq.Array(actor.Movies)).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("ошибка при добавлении актера в базу данных: %w", err)
	}

	return actor.Id, nil
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

func updateActorInDB(db *sql.DB, actorID string, actor Actor) error {
	query := `UPDATE actors SET name = $1, gender = $2, birth_date = $3, movies = $4 WHERE id = $5`
	_, err := db.Exec(query, actor.Name.FirstName, actor.Gender, actor.BirthDate, pq.Array(actor.Movies), actorID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении актера в базе данных: %w", err)
	}
	return nil
}

func deleteActorFromDB(db *sql.DB, actorID string) error {
	if actorID == "" {
		return errors.New("actor_id is required")
	}

	query := `DELETE FROM actors WHERE id = $1`
	result, err := db.Exec(query, actorID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении актера из базы данных: %w", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка при проверке количества удаленных строк: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
