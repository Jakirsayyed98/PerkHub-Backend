package model

import (
	"PerkHub/request"
	"database/sql"
	"time"
)

type Category struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Image           string    `json:"image"`
	Status          bool      `json:"status"`
	HomepageVisible bool      `json:"homepage_visible"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewCategorymodel() Category {
	return Category{}
}

func InsertCategory(db *sql.DB, item *request.Category) error {
	query := `
		INSERT INTO miniapp_categories ( name, description, image, status, homepage_visible, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := db.Exec(query, item.Name, item.Description, item.Image, true, false, time.Now(), time.Now())
	return err
}

func DeleteCategoryByID(db *sql.DB, id string) error {
	query := "DELETE FROM miniapp_categories WHERE id = $1;"
	_, err := db.Exec(query, id)
	return err
}

func UpdateCategory(db *sql.DB, item *request.Category) error {
	query := `
		UPDATE miniapp_categories 
SET 
    name = $1, 
    description = $2, 
    image = $3, 
    status = $4, 
    homepage_visible = $5, 
    updated_at = $6
WHERE id = $7`

	_, err := db.Exec(query, item.Name, item.Description, item.Image, item.Status, item.HomepageVisible, time.Now(), item.ID)
	return err
}

func GetAllCategory(db *sql.DB) ([]*Category, error) {

	query := "SELECT id, name, description, image, status, homepage_visible, created_at, updated_at FROM miniapp_categories"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category []*Category

	for rows.Next() {
		categorym := NewCategorymodel()

		err := rows.Scan(
			&categorym.ID,
			&categorym.Name,
			&categorym.Description,
			&categorym.Image,
			&categorym.Status,
			&categorym.HomepageVisible,
			&categorym.CreatedAt,
			&categorym.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		category = append(category, &categorym)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return category, nil
}

func GetAllHomePageActive(db *sql.DB) ([]*Category, error) {

	query := "SELECT id, name, description, image, status, homepage_visible, created_at, updated_at FROM miniapp_categories WHERE homepage_visible='1' AND status='1'"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var category []*Category

	for rows.Next() {
		categorym := NewCategorymodel()

		err := rows.Scan(
			&categorym.ID,
			&categorym.Name,
			&categorym.Description,
			&categorym.Image,
			&categorym.Status,
			&categorym.HomepageVisible,
			&categorym.CreatedAt,
			&categorym.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		category = append(category, &categorym)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return category, nil
}
