package model

import (
	"PerkHub/request"
	"database/sql"
	"time"
)

type Banner struct {
	ID        string    `json:"id"`         // Assuming ID is a string (could also be UUID)
	Name      string    `json:"name"`       // Name of the item
	BannerId  string    `json:"banner_id"`  // Name of the item
	Image     string    `json:"image"`      // URL or path to the item's image
	Url       string    `json:"url"`        // URL or path to the item's url
	Status    string    `json:"status"`     // Status of the item (e.g., active, inactive)
	CreatedAt time.Time `json:"created_at"` // Timestamp when the item was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the item was last updated
}

func NewBanner() *Banner {
	return &Banner{}
}

func InsertBanner(db *sql.DB, item *request.Banner) error {
	query := `
		INSERT INTO banner_data ( name, banner_id, image, url, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := db.Exec(query, item.Name, item.BannerId, item.Image, item.Url, "1", time.Now(), time.Now())
	return err
}

func UpdateBanner(db *sql.DB, item *request.Banner) error {
	query := `
		UPDATE banner_data 
SET 
    name = $1, 
    status = $2, 
	url = $3,
    updated_at = $4
WHERE id = $5`

	_, err := db.Exec(query, item.Name, item.Status, item.Url, time.Now(), item.ID)
	return err
}

func DeleteBanner(db *sql.DB, id string) error {
	query := "DELETE FROM banner_data WHERE id = $1;"
	_, err := db.Exec(query, id)
	return err
}

func GetAllBanners(db *sql.DB) ([]*Banner, error) {
	query := "SELECT id, name, banner_id, image, url, status, created_at, updated_at FROM banner_data"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []*Banner

	for rows.Next() {
		banner := &Banner{}

		err := rows.Scan(
			&banner.ID,
			&banner.Name,
			&banner.BannerId,
			&banner.Image,
			&banner.Url,
			&banner.Status,
			&banner.CreatedAt,
			&banner.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		banners = append(banners, banner)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return banners, nil
}

func GetBannerbyId(db *sql.DB, id string) ([]*Banner, error) {
	query := "SELECT id, name, banner_id, image, url, status, created_at, updated_at FROM banner_data WHERE banner_id=$1"

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []*Banner

	for rows.Next() {
		banner := &Banner{}

		err := rows.Scan(
			&banner.ID,
			&banner.Name,
			&banner.BannerId,
			&banner.Image,
			&banner.Url,
			&banner.Status,
			&banner.CreatedAt,
			&banner.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		banners = append(banners, banner)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return banners, nil
}
