package model

import (
	"PerkHub/request"
	"PerkHub/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
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
	Status               bool      `db:"status" json:"status"`                                 // Status: '0' for inactive, '1' for active
	UrlType              string    `db:"url_type" json:"url_type"`                             // Type of URL
	CBActive             bool      `db:"cb_active" json:"cb_active"`                           // Cashback active status
	CBPercentage         string    `db:"cb_percentage" json:"cb_percentage"`                   // Cashback percentage
	Url                  string    `db:"url" json:"url"`                                       // URL of the miniapp
	Label                string    `db:"label" json:"label"`                                   // Label for the miniapp
	Banner               string    `db:"banner" json:"banner"`                                 // Banner URL
	Logo                 string    `db:"logo" json:"logo"`                                     // Logo URL
	MacroPublisher       string    `db:"macro_publisher" json:"macro_publisher"`               // Publisher name
	Popular              bool      `db:"popular" json:"popular"`                               // Popular status
	Trending             bool      `db:"trending" json:"trending"`                             // Trending status
	TopCashback          bool      `db:"top_cashback" json:"top_cashback"`                     // Top cashback status
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
func UpdateMiniAppData(db *sql.DB, update *request.MiniAppRequest) error {
	var clauses []string
	var params []interface{}
	paramIndex := 1 // Start the parameter index from 1

	// Check if any field is being updated and add it to the clauses
	if update.MiniAppCategoryID != "" {
		clauses = append(clauses, fmt.Sprintf("miniapp_category_id = $%d", paramIndex))
		params = append(params, update.MiniAppCategoryID)
		paramIndex++
	}
	if update.MiniAppSubcategoryID != "" {
		clauses = append(clauses, fmt.Sprintf("miniapp_subcategory_id = $%d", paramIndex))
		params = append(params, update.MiniAppSubcategoryID)
		paramIndex++
	}
	if update.Name != "" {
		clauses = append(clauses, fmt.Sprintf("name = $%d", paramIndex))
		params = append(params, update.Name)
		paramIndex++
	}
	if update.Icon != "" {
		clauses = append(clauses, fmt.Sprintf("icon = $%d", paramIndex))
		params = append(params, update.Icon)
		paramIndex++
	}
	if update.Description != "" {
		clauses = append(clauses, fmt.Sprintf("description = $%d", paramIndex))
		params = append(params, update.Description)
		paramIndex++
	}
	if update.CashbackTerms != "" {
		clauses = append(clauses, fmt.Sprintf("cashback_terms = $%d", paramIndex))
		params = append(params, update.CashbackTerms)
		paramIndex++
	}
	if update.CashbackRates != "" {
		clauses = append(clauses, fmt.Sprintf("cashback_rates = $%d", paramIndex))
		params = append(params, update.CashbackRates)
		paramIndex++
	}
	if update.UrlType != "" {
		clauses = append(clauses, fmt.Sprintf("url_type = $%d", paramIndex))
		params = append(params, update.UrlType)
		paramIndex++
	}
	if update.CBPercentage != "" {
		clauses = append(clauses, fmt.Sprintf("cb_percentage = $%d", paramIndex))
		params = append(params, update.CBPercentage)
		paramIndex++
	}
	if update.Url != "" {
		clauses = append(clauses, fmt.Sprintf("url = $%d", paramIndex))
		params = append(params, update.Url)
		paramIndex++
	}
	if update.Label != "" {
		clauses = append(clauses, fmt.Sprintf("label = $%d", paramIndex))
		params = append(params, update.Label)
		paramIndex++
	}
	if update.MacroPublisher != "" {
		clauses = append(clauses, fmt.Sprintf("macro_publisher = $%d", paramIndex))
		params = append(params, update.MacroPublisher)
		paramIndex++
	}
	if update.Banner != "" {
		clauses = append(clauses, fmt.Sprintf("banner = $%d", paramIndex))
		params = append(params, update.Banner)
		paramIndex++
	}
	if update.Logo != "" {
		clauses = append(clauses, fmt.Sprintf("logo = $%d", paramIndex))
		params = append(params, update.Logo)
		paramIndex++
	}
	if update.About != "" {
		clauses = append(clauses, fmt.Sprintf("about = $%d", paramIndex))
		params = append(params, update.About)
		paramIndex++
	}
	if update.HowItsWork != "" {
		clauses = append(clauses, fmt.Sprintf("howitswork = $%d", paramIndex))
		params = append(params, update.HowItsWork)
		paramIndex++
	}

	// Add the updated_at timestamp field
	clauses = append(clauses, fmt.Sprintf("updated_at = $%d", paramIndex))
	params = append(params, "NOW()")
	paramIndex++

	// Handle boolean fields properly (Popular, Trending, TopCashback, Status, CBActive)
	if update.Popular {
		clauses = append(clauses, fmt.Sprintf("popular = $%d", paramIndex))
		params = append(params, update.Popular)
		paramIndex++
	} else {
		clauses = append(clauses, fmt.Sprintf("popular = $%d", paramIndex))
		params = append(params, false)
		paramIndex++
	}

	if update.Trending {
		clauses = append(clauses, fmt.Sprintf("trending = $%d", paramIndex))
		params = append(params, update.Trending)
		paramIndex++
	} else {
		clauses = append(clauses, fmt.Sprintf("trending = $%d", paramIndex))
		params = append(params, false)
		paramIndex++
	}

	if update.TopCashback {
		clauses = append(clauses, fmt.Sprintf("top_cashback = $%d", paramIndex))
		params = append(params, update.TopCashback)
		paramIndex++
	} else {
		clauses = append(clauses, fmt.Sprintf("top_cashback = $%d", paramIndex))
		params = append(params, false)
		paramIndex++
	}

	if update.Status {
		clauses = append(clauses, fmt.Sprintf("status = $%d", paramIndex))
		params = append(params, update.Status)
		paramIndex++
	} else {
		clauses = append(clauses, fmt.Sprintf("status = $%d", paramIndex))
		params = append(params, false)
		paramIndex++
	}

	if update.CBActive {
		clauses = append(clauses, fmt.Sprintf("cb_active = $%d", paramIndex))
		params = append(params, update.CBActive)
		paramIndex++
	} else {
		clauses = append(clauses, fmt.Sprintf("cb_active = $%d", paramIndex))
		params = append(params, false)
		paramIndex++
	}

	// Ensure that the MiniApp ID is provided for the update query
	if update.ID == "" {
		return errors.New("missing MiniApp ID for update")
	}

	// Append the WHERE clause at the end
	// clauses = append(clauses, fmt.Sprintf("WHERE id = $%d", paramIndex))
	// params = append(params, update.ID)

	// Construct the final query
	query := "UPDATE miniapp_data SET " + strings.Join(clauses, ", ") + " WHERE id = " + "'" + update.ID + "'"

	// Execute the update query
	_, err := db.Exec(query, params...)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	return nil
}

func ActivateSomekey(db *sql.DB, key, id string, value bool) error {
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
		FROM miniapp_data` // Adjust the table name as per your database schema

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
		miniApp.Icon = utils.ImageUrlGenerator(miniApp.Icon)
		miniApp.Banner = utils.ImageUrlGenerator(miniApp.Banner)
		miniApp.Logo = utils.ImageUrlGenerator(miniApp.Logo)

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
		FROM miniapp_data WHERE popular=true AND status=true`

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

		miniApp.Icon = utils.ImageUrlGenerator(miniApp.Icon)
		miniApp.Banner = utils.ImageUrlGenerator(miniApp.Banner)
		miniApp.Logo = utils.ImageUrlGenerator(miniApp.Logo)

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
		FROM miniapp_data WHERE top_cashback=true AND status=true`

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

		miniApp.Icon = utils.ImageUrlGenerator(miniApp.Icon)
		miniApp.Banner = utils.ImageUrlGenerator(miniApp.Banner)
		miniApp.Logo = utils.ImageUrlGenerator(miniApp.Logo)

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miniApps, nil
}

func GetMiniAppsTrending(db *sql.DB) ([]MiniApp, error) {
	var miniApps []MiniApp
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
		FROM miniapp_data WHERE trending=true AND status=true`

	rows, err := db.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return miniApps, nil
		}
		return nil, err
	}
	defer rows.Close()

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

		miniApp.Icon = utils.ImageUrlGenerator(miniApp.Icon)
		miniApp.Banner = utils.ImageUrlGenerator(miniApp.Banner)
		miniApp.Logo = utils.ImageUrlGenerator(miniApp.Logo)

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
		WHERE name ILIKE $1 AND status=true` // Use ILIKE for case-insensitive matching

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

		miniApp.Icon = utils.ImageUrlGenerator(miniApp.Icon)
		miniApp.Banner = utils.ImageUrlGenerator(miniApp.Banner)
		miniApp.Logo = utils.ImageUrlGenerator(miniApp.Logo)

		miniApps = append(miniApps, miniApp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return miniApps, nil
}

func GetStoresByCategory(db *sql.DB, categoryID string) ([]MiniApp, error) {
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
		WHERE miniapp_category_id = $1 AND status=true` // Match by category ID

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

		miniApp.Icon = utils.ImageUrlGenerator(miniApp.Icon)
		miniApp.Banner = utils.ImageUrlGenerator(miniApp.Banner)
		miniApp.Logo = utils.ImageUrlGenerator(miniApp.Logo)

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

func GetStoreByID(db *sql.DB, id string) (*MiniApp, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %v", err)
	}
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
		WHERE id=$1
	`

	row := db.QueryRow(query, parsedID)

	var miniApp MiniApp
	err = row.Scan(
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
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows // no record found
		}
		return nil, err
	}

	return &miniApp, nil
}
