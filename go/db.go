package swagger

import (
	"database/sql"
	"fmt"
)

func getActorFromDB(db *sql.DB, actorID string) (*Actor, error) {
	query := `SELECT id, name, gender, birth_date FROM actors WHERE id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	var actor Actor
	err = stmt.QueryRow(actorID).Scan(&actor.Id, &actor.Name, &actor.Gender, &actor.BirthDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Возвращаем nil, nil, если актер не найден
		}
		return nil, fmt.Errorf("failed to query actor: %w", err)
	}
	return &actor, nil
}
