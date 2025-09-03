package model

import (
	"PerkHub/request"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type BannerCategory struct {
	ID        string    `json:"id"`         // Assuming ID is a string (could also be UUID)
	Title     string    `json:"title"`      // Name of the item
	Status    bool      `json:"status"`     // Status of the item (e.g., active, inactive)
	CreatedAt time.Time `json:"created_at"` // Timestamp when the item was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the item was last updated
	Banner    []*Banner `json:"banner"`
}

type Banner struct {
	ID               string    `json:"id"`                 // Assuming ID is a string (could also be UUID)
	Name             string    `json:"name"`               // Name of the item
	BannerCategoryId string    `json:"banner_category_id"` // Name of the item
	Image            string    `json:"image"`              // URL or path to the item's image
	Url              string    `json:"url"`                // URL or path to the item's url
	StartDate        string    `json:"start_date"`         // URL or path to the item's url
	EndDate          string    `json:"end_date"`           // URL or path to the item's url
	Status           bool      `json:"status"`             // Status of the item (e.g., active, inactive)
	CreatedAt        time.Time `json:"created_at"`         // Timestamp when the item was created
	UpdatedAt        time.Time `json:"updated_at"`         // Timestamp when the item was last updated
}

func NewBanner() *Banner {
	return &Banner{}
}

func InsertBannerCategory(db *sql.DB, item *request.BannerCategory) error {
	query := `
		INSERT INTO Banner_Category ( title, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := db.Exec(query, item.Title, true, time.Now(), time.Now())
	return err
}

func GetAllBannersCategory(db *sql.DB) ([]*BannerCategory, error) {
	query := "SELECT id, title, status, created_at, updated_at FROM banner_category"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var banners []*BannerCategory

	for rows.Next() {
		banner := &BannerCategory{}

		err := rows.Scan(
			&banner.ID,
			&banner.Title,
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

func InsertBanner(db *sql.DB, item *request.Banner) error {

	query := `
		INSERT INTO banner_data ( name, banner_category_id, image, url, status,start_date,end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9)
	`
	_, err := db.Exec(query, item.Name, item.BannerCategoryId, item.Image, item.Url, true, item.StartDate, item.EndDate, time.Now(), time.Now())
	return err
}

func UpdateBannerData(db *sql.DB, item *request.Banner) error {
	var clauses []string
	var params []interface{}
	paramIndex := 1 // Start the parameter index from 1

	// Check if any field is being updated and add it to the clauses
	if item.Name != "" {
		clauses = append(clauses, fmt.Sprintf("name = $%d", paramIndex))
		params = append(params, item.Name)
		paramIndex++
	}

	clauses = append(clauses, fmt.Sprintf("status = $%d", paramIndex))
	params = append(params, item.Status)
	paramIndex++

	if item.Url != "" {
		clauses = append(clauses, fmt.Sprintf("url = $%d", paramIndex))
		params = append(params, item.Url)
		paramIndex++
	}
	if item.StartDate != "" {
		clauses = append(clauses, fmt.Sprintf("start_date = $%d", paramIndex))
		params = append(params, item.StartDate)
		paramIndex++
	}
	if item.EndDate != "" {
		clauses = append(clauses, fmt.Sprintf("end_date = $%d", paramIndex))
		params = append(params, item.EndDate)
		paramIndex++
	}

	if item.Image != "" {
		clauses = append(clauses, fmt.Sprintf("image = $%d", paramIndex))
		params = append(params, item.Image)
		paramIndex++
	}

	// Add the updated_at timestamp field
	clauses = append(clauses, fmt.Sprintf("updated_at = $%d", paramIndex))
	params = append(params, time.Now())
	paramIndex++

	// Ensure that the Banner ID is provided for the update query
	if item.ID == "" {
		return errors.New("missing Banner ID for update")
	}

	// Construct the final query
	query := "UPDATE banner_data SET " + strings.Join(clauses, ", ") + " WHERE id =" + "'" + item.ID + "'"
	// params = append(params, item.ID) // Add ID at the end for WHERE clause

	// Execute the update query
	_, err := db.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	return nil
}

func DeleteBanner(db *sql.DB, id string) error {
	query := "DELETE FROM banner_data WHERE id = $1;"
	_, err := db.Exec(query, id)
	return err
}
func GetBannersByCategoryID(db *sql.DB, categoryID string) ([]*Banner, error) {
	// Modify the query to filter by banner_category_id
	query := "SELECT id, name, banner_category_id, image, url, status, start_date, end_date, created_at, updated_at FROM banner_data WHERE banner_category_id = $1"

	// Execute the query with the provided categoryID as a parameter
	rows, err := db.Query(query, categoryID)
	if err != nil {
		return nil, errors.New("no data found")
	}
	defer rows.Close()

	// Initialize a slice to hold the banners
	var banners []*Banner

	// Loop through the result set and scan the data into Banner structs
	for rows.Next() {
		banner := &Banner{}
		err := rows.Scan(
			&banner.ID,
			&banner.Name,
			&banner.BannerCategoryId,
			&banner.Image,
			&banner.Url,
			&banner.Status,
			&banner.StartDate,
			&banner.EndDate,
			&banner.CreatedAt,
			&banner.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		banners = append(banners, banner)
	}

	// Check if there were any errors while scanning rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Return the list of banners for the specified category
	return banners, nil
}

func GetBannerbyId(db *sql.DB, id string) (*Banner, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %v", err)
	}

	query := "SELECT id,name, banner_category_id, image, url, status, start_date, end_date, created_at, updated_at FROM banner_data WHERE id = $1"
	// Execute the query with the provided categoryID as a parameter
	row := db.QueryRow(query, parsedID)
	if err != nil {
		return nil, err
	}

	// Initialize a slice to hold the banners

	banner := &Banner{}
	err = row.Scan(
		&banner.ID,
		&banner.Name,
		&banner.BannerCategoryId,
		&banner.Image,
		&banner.Url,
		&banner.Status,
		&banner.StartDate,
		&banner.EndDate,
		&banner.CreatedAt,
		&banner.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	// Check if there were any errors while scanning rows
	if err := row.Err(); err != nil {
		return nil, err
	}

	// Return the list of banners for the specified category
	return banner, nil
}
