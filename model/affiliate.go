package model

import (
	"PerkHub/request"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Affiliate struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	URL       string    `json:"url"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewAffiliate() *Affiliate {
	return &Affiliate{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func CreateAffiliate(db *sql.DB, affiliate *request.CreateAffiliateRequest) error {
	query := "INSERT INTO affiliates (name, key, url, status, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW(), NOW())"
	_, err := db.Exec(query, affiliate.Name, affiliate.Key, affiliate.URL, true)
	return err
}

func UpdateAffiliate(db *sql.DB, affiliate *request.CreateAffiliateRequest, id uuid.UUID) error {
	query := "UPDATE affiliates SET name = $1, key = $2, url = $3, status = $4, updated_at = NOW() WHERE id = $5"
	_, err := db.Exec(query, affiliate.Name, affiliate.Key, affiliate.URL, affiliate.Status, id)
	return err
}

func DeleteAffiliate(db *sql.DB, id uuid.UUID) error {
	query := "DELETE FROM affiliates WHERE id = $1"
	_, err := db.Exec(query, id)
	return err
}

func GetAffiliate(db *sql.DB, id uuid.UUID) (*Affiliate, error) {
	query := "SELECT id, name, key, url, status, created_at, updated_at FROM affiliates WHERE id = $1"
	row := db.QueryRow(query, id)

	affiliate := &Affiliate{}
	err := row.Scan(&affiliate.ID, &affiliate.Name, &affiliate.Key, &affiliate.URL, &affiliate.Status, &affiliate.CreatedAt, &affiliate.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return affiliate, nil
}

func UpdateAffiliateFlag(db *sql.DB, id uuid.UUID, status bool) error {
	query := "UPDATE affiliates SET status = $1, updated_at = NOW() WHERE id = $2"
	_, err := db.Exec(query, status, id)
	return err
}

func ListAffiliates(db *sql.DB) ([]*Affiliate, error) {
	query := "SELECT id, name, key, url, status, created_at, updated_at FROM affiliates"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var affiliates []*Affiliate
	for rows.Next() {
		affiliate := &Affiliate{}
		err := rows.Scan(&affiliate.ID, &affiliate.Name, &affiliate.Key, &affiliate.URL, &affiliate.Status, &affiliate.CreatedAt, &affiliate.UpdatedAt)
		if err != nil {
			return nil, err
		}
		affiliates = append(affiliates, affiliate)
	}
	return affiliates, nil
}
