package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Offer struct {
	ID                uuid.UUID `db:"id" json:"id"`
	OfferID           int64     `db:"offer_id" json:"offer_id"`
	StoreID           uuid.UUID `db:"store_id" json:"store_id"`
	StoreName         string    `db:"store_name" json:"store_name"`
	Title             string    `db:"title" json:"title"`
	Description       string    `db:"description" json:"description"`
	TermsAndCondition string    `db:"terms_and_condition" json:"terms_and_condition"`
	CouponCode        string    `db:"coupon_code" json:"coupon_code"`
	Image             string    `db:"image" json:"image"`
	Type              string    `db:"type" json:"type"` // "coupon" or "offer"
	Status            bool      `db:"status" json:"status"`
	URL               string    `db:"url" json:"url"`
	StartDate         string    `db:"start_date" json:"start_date,omitempty"`
	EndDate           string    `db:"end_date" json:"end_date,omitempty"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

func NewOffers() *Offer {
	return &Offer{}
}

func InsertOffers(db *sql.DB, offer *Offer) error {
	query := `
		INSERT INTO offers (
			offer_id, store_id, store_name, title, description,
			terms_and_condition, coupon_code, image, type, status,
			url, start_date, end_date, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15
		)`

	err := db.QueryRow(
		query,
		offer.OfferID,
		offer.StoreID,
		offer.StoreName,
		offer.Title,
		offer.Description,
		offer.TermsAndCondition,
		offer.CouponCode,
		offer.Image,
		offer.Type,
		offer.Status,
		offer.URL,
		offer.StartDate,
		offer.EndDate,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err.Err()
	}
	return nil
}

func GetSearchOffers(db *sql.DB, storeName string) ([]Offer, error) {
	query := `SELECT id, offer_id, store_id, store_name, title, description,
		       terms_and_condition, coupon_code, image, type, status,
		       url, start_date, end_date, created_at, updated_at
		FROM offers
		WHERE status = true
		  AND store_name = $1
		  AND start_date <= CURRENT_DATE
		  AND end_date >= CURRENT_DATE`

	rows, err := db.Query(query, storeName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []Offer

	for rows.Next() {
		var offer Offer
		err := rows.Scan(
			&offer.ID,
			&offer.OfferID,
			&offer.StoreID,
			&offer.StoreName,
			&offer.Title,
			&offer.Description,
			&offer.TermsAndCondition,
			&offer.CouponCode,
			&offer.Image,
			&offer.Type,
			&offer.Status,
			&offer.URL,
			&offer.StartDate,
			&offer.EndDate,
			&offer.CreatedAt,
			&offer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		offers = append(offers, offer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return offers, nil
}

func GetAllOfferList(db *sql.DB) ([]Offer, error) {
	query := `SELECT id, offer_id, store_id, store_name, title, description,
		       terms_and_condition, coupon_code, image, type, status,
		       url, start_date, end_date, created_at, updated_at
		FROM offers
		WHERE status = true
		  AND store_name = $1
		  AND start_date <= CURRENT_DATE
		  AND end_date >= CURRENT_DATE`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var offers []Offer

	for rows.Next() {
		var offer Offer
		err := rows.Scan(
			&offer.ID,
			&offer.OfferID,
			&offer.StoreID,
			&offer.StoreName,
			&offer.Title,
			&offer.Description,
			&offer.TermsAndCondition,
			&offer.CouponCode,
			&offer.Image,
			&offer.Type,
			&offer.Status,
			&offer.URL,
			&offer.StartDate,
			&offer.EndDate,
			&offer.CreatedAt,
			&offer.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		offers = append(offers, offer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return offers, nil
}

func OfferExists(db *sql.DB, offerId string) (bool, error) {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM offers WHERE offer_id=$1)", offerId).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check existence: %w", err)
	}
	return exists, nil
}
