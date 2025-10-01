package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type GameCategory struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Icon   string    `json:"icon"`
	Status bool      `json:"status"`
}

func NewGameCategory() *GameCategory {
	return &GameCategory{}
}

func InsertGamesCategory(db *sql.DB, name, icon string) error {
	sqlQuery := `INSERT INTO game_categories ( name, icon, created_at, updated_at )
	 VALUES ($1,$2,$3,$4)`

	_, err := db.Exec(sqlQuery, name, icon, time.Now(), time.Now())

	return err
}

func (s *GameCategory) FindGameCategoryByNameOrId(db *sql.DB, id, name string) (*GameCategory, error) {
	query := `SELECT id, name, icon, status FROM game_categories WHERE name = $1;`
	row := db.QueryRow(query, name)

	var category GameCategory

	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.Icon,
		&category.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return &category, nil
}

func GetGameCategories(db *sql.DB) ([]GameCategory, error) {
	query := `SELECT id, name, icon, status FROM game_categories WHERE status = '1';`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var categories []GameCategory

	for rows.Next() {
		var category GameCategory
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Icon,
			&category.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return categories, nil
}

func AdminGetGameCategories(db *sql.DB) ([]GameCategory, error) {
	query := `SELECT id, name, icon, status FROM game_categories;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var categories []GameCategory

	for rows.Next() {
		var category GameCategory
		err := rows.Scan(
			&category.Id,
			&category.Name,
			&category.Icon,
			&category.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return categories, nil
}

func ActivateDeactiveGameCategoryKey(db *sql.DB, id string, value bool) error {
	query := `UPDATE game_categories SET status = $1 WHERE id = $2`
	_, err := db.Exec(query, value, id)
	return err
}
