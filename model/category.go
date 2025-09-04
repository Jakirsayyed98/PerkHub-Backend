package model

import (
	"PerkHub/request"
	"PerkHub/utils"
	"database/sql"
	"fmt"
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

		categorym.Image = utils.ImageUrlGenerator(categorym.Image)
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

func GetCategoryByID(db *sql.DB, id string) (*Category, error) {
	query := "SELECT id, name, description, image, status, homepage_visible, created_at, updated_at FROM miniapp_categories WHERE id = $1"
	row := db.QueryRow(query, id)

	var category Category
	err := row.Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.Image,
		&category.Status,
		&category.HomepageVisible,
		&category.CreatedAt,
		&category.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	// category.Image = utils.ImageUrlGenerator(category.Image)
	return &category, nil
}

func ActivateDeactiveCategorykey(db *sql.DB, key, id string, value bool) error {
	query := fmt.Sprintf("UPDATE miniapp_data SET %s = $1 WHERE id = $2", key)
	if _, err := db.Exec(query, value, id); err != nil {
		return err
	}
	return nil
}

func CategoryExists(db *sql.DB, name string) (bool, error) {
	query := "SELECT 1 FROM miniapp_categories WHERE name = $1 LIMIT 1"
	row := db.QueryRow(query, name)

	var exists int
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil // category does not exist
		}
		return false, err // some other error
	}

	return true, nil // category exists
}

func CategoryByName(db *sql.DB, name string) (string, error) {
	query := "SELECT id FROM miniapp_categories WHERE name = $1 LIMIT 1"
	row := db.QueryRow(query, name)

	var id string
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // category does not exist
		}
		return "", err // some other error
	}

	return id, nil // category exists
}
