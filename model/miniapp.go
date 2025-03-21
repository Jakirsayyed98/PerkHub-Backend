package model

import (
	"PerkHub/request"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MiniApp struct {
	ID                   uuid.UUID `db:"id" json:"id"`                                         // Unique identifier
	MiniAppCategoryID    string    `db:"miniapp_category_id" json:"miniapp_category_id"`       // Category ID
	MiniAppSubcategoryID string    `db:"miniapp_subcategory_id" json:"miniapp_subcategory_id"` // Subcategory ID
	Name                 string    `db:"name" json:"name"`                                     // Name of the miniapp
	Icon                 string    `db:"icon" json:"icon"`                                     // URL or path to the icon
	Description          string    `db:"description" json:"description"`                       // Description of the miniapp
	CashbackTerms        string    `db:"cashback_terms" json:"cashback_terms"`                 // Terms for cashback
	CashbackRates        string    `db:"cashback_rates" json:"cashback_rates"`                 // Rates for cashback
	Status               string    `db:"status" json:"status"`                                 // Status: '0' for inactive, '1' for active
	UrlType              string    `db:"url_type" json:"url_type"`                             // Type of URL
	CBActive             string    `db:"cb_active" json:"cb_active"`                           // Cashback active status
	CBPercentage         string    `db:"cb_percentage" json:"cb_percentage"`                   // Cashback percentage
	Url                  string    `db:"url" json:"url"`                                       // URL of the miniapp
	Label                string    `db:"label" json:"label"`                                   // Label for the miniapp
	Banner               string    `db:"banner" json:"banner"`                                 // Banner URL
	Logo                 string    `db:"logo" json:"logo"`                                     // Logo URL
	MacroPublisher       string    `db:"macro_publisher" json:"macro_publisher"`               // Publisher name
	Popular              string    `db:"popular" json:"popular"`                               // Popular status
	Trending             string    `db:"trending" json:"trending"`                             // Trending status
	TopCashback          string    `db:"top_cashback" json:"top_cashback"`                     // Top cashback status
	About                string    `db:"about" json:"about"`                                   // About information
	HowItsWork           string    `db:"howitswork" json:"howitswork"`                         // How it works information
	CreatedAt            time.Time `db:"created_at" json:"created_at"`                         // Creation timestamp
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
}

type GenerateSubId struct {
	ID        uuid.UUID `db:"id" json:"id"`
	StoreID   string    `db:"miniapp_id" json:"miniapp_id"`
	UserID    string    `db:"user_id" json:"user_id"`
	SubID1    string    `db:"subid1" json:"subid1"`
	SubID2    string    `db:"subid2" json:"subid2"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func InsertGenratedSubId(db *sql.DB, miniapp_id, userId, subid1, subid2 string) error {

	query := `
    INSERT INTO genratedsubid_data (
        miniapp_id,user_id,subid1,subid2,created_at,updated_at
    ) VALUES (
        $1, $2, $3, $4, $5, $6
    )`

	_, err := db.Exec(query,
		miniapp_id, userId, subid1, subid2,
		time.Now(),
		time.Now(),
	)

	return err
}

func InsertMiniAppData(db *sql.DB, req *request.MiniAppRequest) error {

	query := `
    INSERT INTO miniapp_data (
        miniapp_category_id,
        miniapp_subcategory_id,
        name,
        icon,
        description,
        cashback_terms,
        cashback_rates,
        status,
        url_type,
        cb_active,
        cb_percentage,
        url,
        label,
        banner,
        logo,
        macro_publisher,
        popular,
        trending,
        top_cashback,
        about,
        howitswork,
        created_at,
        updated_at
    ) VALUES (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,
        $14, $15, $16, $17, $18, $19, $20, $21, $22, $23
    )`

	_, err := db.Exec(query,
		req.MiniAppCategoryID,
		req.MiniAppSubcategoryID,
		req.Name,
		req.Icon,
		req.Description,
		req.CashbackTerms,
		req.CashbackRates,
		req.Status,
		req.UrlType,
		req.CBActive,
		req.CBPercentage,
		req.Url,
		req.Label,
		req.Banner,
		req.Logo,
		req.MacroPublisher,
		req.Popular,
		req.Trending,
		req.TopCashback,
		req.About,
		req.HowItsWork,
		req.CreatedAt,
		req.UpdatedAt,
	)

	return err
}

func ActivateSomekey(db *sql.DB, key, id, value string) error {
	query := fmt.Sprintf("UPDATE miniapp_data SET %s = $1 WHERE id = $2", key)
	if _, err := db.Exec(query, value, id); err != nil {
		return err
	}
	return nil
}

func GetAllMiniApps(db *sql.DB) ([]MiniApp, error) {
	query := `
		SELECT 
			id, 
			miniapp_category_id, 
			miniapp_subcategory_id, 
			name, 
			icon, 
			description, 
			cashback_terms, 
			cashback_rates, 
			status, 
			url_type, 
			cb_active, 
			cb_percentage, 
			url, 
			label, 
			banner, 
			logo, 
			macro_publisher, 
			popular, 
			trending, 
			top_cashback, 
			about, 
			howitswork, 
			created_at, 
			updated_at 
		FROM miniapp_data WHERE status='1'` // Adjust the table name as per your database schema

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var miniApps []MiniApp

	for rows.Next() {
		var miniApp MiniApp

		err := rows.Scan(
			&miniApp.ID,
			&miniApp.MiniAppCategoryID,
			&miniApp.MiniAppSubcategoryID,
			&miniApp.Name,
			&miniApp.Icon,
			&miniApp.Description,
			&miniApp.CashbackTerms,
			&miniApp.CashbackRates,
			&miniApp.Status,
			&miniApp.UrlType,
			&miniApp.CBActive,
			&miniApp.CBPercentage,
			&miniApp.Url,
			&miniApp.Label,
			&miniApp.Banner,
			&miniApp.Logo,
			&miniApp.MacroPublisher,
			&miniApp.Popular,
			&miniApp.Trending,
			&miniApp.TopCashback,
			&miniApp.About,
			&miniApp.HowItsWork,
			&miniApp.CreatedAt,
			&miniApp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return miniApps, nil
}

func GetMiniAppsPopular(db *sql.DB) ([]MiniApp, error) {
	query := `
		SELECT 
			id, 
			miniapp_category_id, 
			miniapp_subcategory_id, 
			name, 
			icon, 
			description, 
			cashback_terms, 
			cashback_rates, 
			status, 
			url_type, 
			cb_active, 
			cb_percentage, 
			url, 
			label, 
			banner, 
			logo, 
			macro_publisher, 
			popular, 
			trending, 
			top_cashback, 
			about, 
			howitswork, 
			created_at, 
			updated_at 
		FROM miniapp_data WHERE popular='1' AND status='1'`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var miniApps []MiniApp

	for rows.Next() {
		var miniApp MiniApp
		err := rows.Scan(
			&miniApp.ID,
			&miniApp.MiniAppCategoryID,
			&miniApp.MiniAppSubcategoryID,
			&miniApp.Name,
			&miniApp.Icon,
			&miniApp.Description,
			&miniApp.CashbackTerms,
			&miniApp.CashbackRates,
			&miniApp.Status,
			&miniApp.UrlType,
			&miniApp.CBActive,
			&miniApp.CBPercentage,
			&miniApp.Url,
			&miniApp.Label,
			&miniApp.Banner,
			&miniApp.Logo,
			&miniApp.MacroPublisher,
			&miniApp.Popular,
			&miniApp.Trending,
			&miniApp.TopCashback,
			&miniApp.About,
			&miniApp.HowItsWork,
			&miniApp.CreatedAt,
			&miniApp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miniApps, nil
}

func GetMiniAppsTopCashback(db *sql.DB) ([]MiniApp, error) {
	query := `
		SELECT 
			id, 
			miniapp_category_id, 
			miniapp_subcategory_id, 
			name, 
			icon, 
			description, 
			cashback_terms, 
			cashback_rates, 
			status, 
			url_type, 
			cb_active, 
			cb_percentage, 
			url, 
			label, 
			banner, 
			logo, 
			macro_publisher, 
			popular, 
			trending, 
			top_cashback, 
			about, 
			howitswork, 
			created_at, 
			updated_at 
		FROM miniapp_data WHERE top_cashback='1' AND status='1'`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var miniApps []MiniApp

	for rows.Next() {
		var miniApp MiniApp
		err := rows.Scan(
			&miniApp.ID,
			&miniApp.MiniAppCategoryID,
			&miniApp.MiniAppSubcategoryID,
			&miniApp.Name,
			&miniApp.Icon,
			&miniApp.Description,
			&miniApp.CashbackTerms,
			&miniApp.CashbackRates,
			&miniApp.Status,
			&miniApp.UrlType,
			&miniApp.CBActive,
			&miniApp.CBPercentage,
			&miniApp.Url,
			&miniApp.Label,
			&miniApp.Banner,
			&miniApp.Logo,
			&miniApp.MacroPublisher,
			&miniApp.Popular,
			&miniApp.Trending,
			&miniApp.TopCashback,
			&miniApp.About,
			&miniApp.HowItsWork,
			&miniApp.CreatedAt,
			&miniApp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miniApps, nil
}

func GetMiniAppsTrending(db *sql.DB) ([]MiniApp, error) {
	query := `
		SELECT 
			id, 
			miniapp_category_id, 
			miniapp_subcategory_id, 
			name, 
			icon, 
			description, 
			cashback_terms, 
			cashback_rates, 
			status, 
			url_type, 
			cb_active, 
			cb_percentage, 
			url, 
			label, 
			banner, 
			logo, 
			macro_publisher, 
			popular, 
			trending, 
			top_cashback, 
			about, 
			howitswork, 
			created_at, 
			updated_at 
		FROM miniapp_data WHERE trending='1' AND status='1'`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var miniApps []MiniApp

	for rows.Next() {
		var miniApp MiniApp
		err := rows.Scan(
			&miniApp.ID,
			&miniApp.MiniAppCategoryID,
			&miniApp.MiniAppSubcategoryID,
			&miniApp.Name,
			&miniApp.Icon,
			&miniApp.Description,
			&miniApp.CashbackTerms,
			&miniApp.CashbackRates,
			&miniApp.Status,
			&miniApp.UrlType,
			&miniApp.CBActive,
			&miniApp.CBPercentage,
			&miniApp.Url,
			&miniApp.Label,
			&miniApp.Banner,
			&miniApp.Logo,
			&miniApp.MacroPublisher,
			&miniApp.Popular,
			&miniApp.Trending,
			&miniApp.TopCashback,
			&miniApp.About,
			&miniApp.HowItsWork,
			&miniApp.CreatedAt,
			&miniApp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miniApps, nil
}

func SearchMiniApps(db *sql.DB, name string) ([]MiniApp, error) {
	// Prepare the query to search for mini apps by name
	query := `
		SELECT 
			id, 
			miniapp_category_id, 
			miniapp_subcategory_id, 
			name, 
			icon, 
			description, 
			cashback_terms, 
			cashback_rates, 
			status, 
			url_type, 
			cb_active, 
			cb_percentage, 
			url, 
			label, 
			banner, 
			logo, 
			macro_publisher, 
			popular, 
			trending, 
			top_cashback, 
			about, 
			howitswork, 
			created_at, 
			updated_at 
		FROM miniapp_data
		WHERE name ILIKE $1 AND status='1'` // Use ILIKE for case-insensitive matching

	rows, err := db.Query(query, "%"+name+"%") // Use wildcards for searching
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var miniApps []MiniApp

	for rows.Next() {
		var miniApp MiniApp

		err := rows.Scan(
			&miniApp.ID,
			&miniApp.MiniAppCategoryID,
			&miniApp.MiniAppSubcategoryID,
			&miniApp.Name,
			&miniApp.Icon,
			&miniApp.Description,
			&miniApp.CashbackTerms,
			&miniApp.CashbackRates,
			&miniApp.Status,
			&miniApp.UrlType,
			&miniApp.CBActive,
			&miniApp.CBPercentage,
			&miniApp.Url,
			&miniApp.Label,
			&miniApp.Banner,
			&miniApp.Logo,
			&miniApp.MacroPublisher,
			&miniApp.Popular,
			&miniApp.Trending,
			&miniApp.TopCashback,
			&miniApp.About,
			&miniApp.HowItsWork,
			&miniApp.CreatedAt,
			&miniApp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miniApps, nil
}

func GetMiniAppsByCategoryID(db *sql.DB, categoryID string) ([]MiniApp, error) {
	// Prepare the query to search for mini apps by category ID
	query := `
		SELECT 
			id, 
			miniapp_category_id, 
			miniapp_subcategory_id, 
			name, 
			icon, 
			description, 
			cashback_terms, 
			cashback_rates, 
			status, 
			url_type, 
			cb_active, 
			cb_percentage, 
			url, 
			label, 
			banner, 
			logo, 
			macro_publisher, 
			popular, 
			trending, 
			top_cashback, 
			about, 
			howitswork, 
			created_at, 
			updated_at 
		FROM miniapp_data
		WHERE miniapp_category_id = $1 AND status='1'` // Match by category ID

	rows, err := db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var miniApps []MiniApp

	for rows.Next() {
		var miniApp MiniApp

		if err := rows.Scan(
			&miniApp.ID,
			&miniApp.MiniAppCategoryID,
			&miniApp.MiniAppSubcategoryID,
			&miniApp.Name,
			&miniApp.Icon,
			&miniApp.Description,
			&miniApp.CashbackTerms,
			&miniApp.CashbackRates,
			&miniApp.Status,
			&miniApp.UrlType,
			&miniApp.CBActive,
			&miniApp.CBPercentage,
			&miniApp.Url,
			&miniApp.Label,
			&miniApp.Banner,
			&miniApp.Logo,
			&miniApp.MacroPublisher,
			&miniApp.Popular,
			&miniApp.Trending,
			&miniApp.TopCashback,
			&miniApp.About,
			&miniApp.HowItsWork,
			&miniApp.CreatedAt,
			&miniApp.UpdatedAt,
		); err != nil {
			return nil, err
		}

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miniApps, nil
}

func DeleteMiniAppByID(db *sql.DB, id string) error {
	query := "DELETE FROM miniapp_data WHERE id = $1;"
	_, err := db.Exec(query, id)
	return err

}
