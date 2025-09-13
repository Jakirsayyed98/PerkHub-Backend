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

type MiniApp struct {
	ID                uuid.UUID `db:"id" json:"id"`
	MiniAppCategoryID uuid.UUID `db:"miniapp_category_id" json:"miniapp_category_id"`
	Name              string    `db:"name" json:"name"`
	Icon              string    `db:"icon" json:"icon"`
	Logo              string    `db:"logo" json:"logo"`
	Banner            string    `db:"banner" json:"banner"`
	Description       string    `db:"description" json:"description"`
	About             string    `db:"about" json:"about"`
	CashbackTerms     string    `db:"cashback_terms" json:"cashback_terms"`
	CBActive          bool      `db:"is_cb_active" json:"is_cb_active"`
	CBPercentage      string    `db:"cb_percentage" json:"cb_percentage"`
	Url               string    `db:"url" json:"url"`
	UrlType           string    `db:"url_type" json:"url_type"`
	MacroPublisher    uuid.UUID `db:"macro_publisher" json:"macro_publisher"`
	Active            bool      `db:"is_active" json:"is_active"`
	Popular           bool      `db:"is_popular" json:"is_popular"`
	Trending          bool      `db:"is_trending" json:"is_trending"`
	TopCashback       bool      `db:"is_top_cashback" json:"is_top_cashback"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
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

func InsertGenratedSubId(db *sql.DB, miniappID, userID, subID1, subID2 string) error {
	query := `
		INSERT INTO genratedsubid_data (miniapp_id, user_id, subid1, subid2, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, miniappID, userID, subID1, subID2, time.Now(), time.Now())
	return err
}

// -------------------- Insert --------------------

func InsertMiniApp(db *sql.DB, req *request.MiniAppRequest) error {
	query := `
	INSERT INTO miniapp_data (
		miniapp_category_id, name, icon, logo, description, about,
		cashback_terms, is_cb_active, cb_percentage, url, url_type,
		macro_publisher, is_active, is_popular, is_trending, is_top_cashback,banner,
		created_at, updated_at
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)
	`
	_, err := db.Exec(query,
		req.MiniAppCategoryID, req.Name, req.Icon, req.Logo, req.Description, req.About,
		req.CashbackTerms, req.CBActive, req.CBPercentage, req.Url, req.UrlType,
		req.MacroPublisher, req.Status, req.Popular, req.Trending, req.TopCashback, req.Banner,
		time.Now(), time.Now(),
	)
	return err
}

// -------------------- Update --------------------

func UpdateMiniApp(db *sql.DB, update *request.MiniAppRequest) error {
	if update.ID == "" {
		return errors.New("missing MiniApp ID for update")
	}

	clauses := []string{}
	params := []interface{}{}
	i := 1

	addField := func(name string, value interface{}) {
		if value != nil {
			clauses = append(clauses, fmt.Sprintf("%s=$%d", name, i))
			params = append(params, value)
			i++
		}
	}

	addField("miniapp_category_id", update.MiniAppCategoryID)
	addField("name", update.Name)
	addField("icon", update.Icon)
	addField("logo", update.Logo)
	addField("description", update.Description)
	addField("about", update.About)
	addField("cashback_terms", update.CashbackTerms)
	addField("is_cb_active", update.CBActive)
	addField("cb_percentage", update.CBPercentage)
	addField("url", update.Url)
	addField("url_type", update.UrlType)
	addField("banner", update.Banner)
	addField("macro_publisher", update.MacroPublisher)
	addField("is_active", update.Status)
	addField("is_popular", update.Popular)
	addField("is_trending", update.Trending)
	addField("is_top_cashback", update.TopCashback)
	addField("updated_at", time.Now())

	query := "UPDATE miniapp_data SET " +
		strings.Join(clauses, ", ") +
		fmt.Sprintf(" WHERE id=$%d", i)
	params = append(params, update.ID)

	_, err := db.Exec(query, params...)
	return err
}

// -------------------- Toggle Boolean --------------------

func ToggleMiniAppFlag(db *sql.DB, field string, id string, value bool) error {
	fmt.Printf("Toggling field %s to %v for MiniApp ID %s\n", field, value, id)
	query := fmt.Sprintf("UPDATE miniapp_data SET %s=$1 WHERE id=$2", field)
	_, err := db.Exec(query, value, id)
	return err
}

func MiniAppExists(db *sql.DB, name string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM miniapp_data WHERE name=$1)", name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}

// -------------------- Scan Helper --------------------

func scanMiniApp(rowScanner interface {
	Scan(dest ...interface{}) error
}) (MiniApp, error) {
	var m MiniApp
	err := rowScanner.Scan(
		&m.ID, &m.MiniAppCategoryID, &m.Name, &m.Icon, &m.Logo, &m.Banner,
		&m.Description, &m.About, &m.CashbackTerms, &m.CBActive, &m.CBPercentage,
		&m.Url, &m.UrlType, &m.MacroPublisher, &m.Active, &m.Popular,
		&m.Trending, &m.TopCashback, &m.CreatedAt, &m.UpdatedAt,
	)
	return m, err
}

// -------------------- Get Helpers --------------------

func GetMiniAppByID(db *sql.DB, id string) (*MiniApp, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	query := `SELECT id, miniapp_category_id, name, icon, logo,banner, description, about,
		cashback_terms, is_cb_active, cb_percentage, url, url_type, macro_publisher,
		is_active, is_popular, is_trending, is_top_cashback, created_at, updated_at
		FROM miniapp_data WHERE id=$1`
	row := db.QueryRow(query, parsedID)
	m, err := scanMiniApp(row)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &m, nil
}

func GetAllMiniApps(db *sql.DB) ([]MiniApp, error) {
	return GetMiniAppsByCondition(db, "")
}

func GetMiniAppsByCondition(db *sql.DB, condition string) ([]MiniApp, error) {
	query := `SELECT id, miniapp_category_id, name, icon, logo,banner, description, about,
		cashback_terms, is_cb_active, cb_percentage, url, url_type, macro_publisher,
		is_active, is_popular, is_trending, is_top_cashback, created_at, updated_at
		FROM miniapp_data ` + condition
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []MiniApp
	for rows.Next() {
		m, err := scanMiniApp(rows)
		if err != nil {
			return nil, err
		}
		apps = append(apps, m)
	}
	return apps, rows.Err()
}

// -------------------- Specific Queries --------------------

func GetMiniAppsPopular(db *sql.DB) ([]MiniApp, error) {
	return GetMiniAppsByCondition(db, "WHERE is_popular=true AND is_active=true")
}

func GetMiniAppsTrending(db *sql.DB) ([]MiniApp, error) {
	return GetMiniAppsByCondition(db, "WHERE is_trending=true AND is_active=true")
}

func GetMiniAppsTopCashback(db *sql.DB) ([]MiniApp, error) {
	return GetMiniAppsByCondition(db, "WHERE is_top_cashback=true AND is_active=true")
}

func SearchMiniApps(db *sql.DB, name string) ([]MiniApp, error) {
	return GetMiniAppsByCondition(db, fmt.Sprintf("WHERE name ILIKE '%%%s%%' AND is_active=true", name))
}

func GetStoresByCategory(db *sql.DB, categoryID string) ([]MiniApp, error) {
	return GetMiniAppsByCondition(db, fmt.Sprintf("WHERE miniapp_category_id='%s' AND is_active=true", categoryID))
}

// -------------------- Delete --------------------

func DeleteMiniAppByID(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM miniapp_data WHERE id=$1", id)
	return err
}
