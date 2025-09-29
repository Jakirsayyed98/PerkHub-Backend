package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Offer struct {
	ID                string    `db:"id" json:"id"`
	OfferID           string    `db:"offer_id" json:"offer_id"`
	StoreID           string    `db:"store_id" json:"store_id"`
	StoreName         string    `db:"store_name" json:"store_name"`
	Title             string    `db:"title" json:"title"`
	Description       string    `db:"description" json:"description"`
	TermsAndCondition string    `db:"terms_and_condition" json:"terms_and_condition"`
	CouponCode        string    `db:"coupon_code" json:"coupon_code"`
	Image             string    `db:"image" json:"image"`
	Type              string    `db:"type" json:"type"` // must be "coupon" or "offer"
	Status            bool      `db:"status" json:"status"`
	URL               string    `db:"url" json:"url"`
	StartDate         string    `db:"start_date" json:"start_date"`
	EndDate           string    `db:"end_date" json:"end_date"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

func NewOffers() *Offer {
	return &Offer{}
}
func InsertOffer(db *sql.DB, offer *Offer) error {
	query := `
		INSERT INTO offers (
			offer_id, store_id, store_name, title, description,
			terms_and_condition, coupon_code, image, type, status,
			url, start_date, end_date, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15
		)
	`

	// Insert
	_, err := db.Exec(
		query,
		offer.OfferID, // varchar
		offer.StoreID, // UUID
		offer.StoreName,
		offer.Title,
		offer.Description,
		offer.TermsAndCondition,
		offer.CouponCode,
		offer.Image,
		offer.Type,
		offer.Status,
		offer.URL,
		offer.StartDate, // string (but should be DATE ideally)
		offer.EndDate,   // string
		time.Now(),      // created_at
		time.Now(),      // updated_at
	)

	return err
}

func GetAllOfferList(db *sql.DB, offerType string) ([]Offer, error) {
	query := `SELECT id, offer_id, store_id, store_name, title, description,
		       terms_and_condition, coupon_code, image, type, status,
		       url, start_date, end_date, created_at, updated_at
		FROM offers
		WHERE status = true AND type = $1`

	rows, err := db.Query(query, offerType)
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

func SearchOffersByStoreName(db *sql.DB, storeName string) ([]Offer, error) {
	query := `SELECT id, offer_id, store_id, store_name, title, description,
		       terms_and_condition, coupon_code, image, type, status,
		       url, start_date, end_date, created_at, updated_at
		FROM offers
		WHERE status = true AND store_name = $1`

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

func GetRandomOffers(db *sql.DB) ([]Offer, error) {
	query := `SELECT id, offer_id, store_id, store_name, title, description,
		       terms_and_condition, coupon_code, image, type, status,
		       url, start_date, end_date, created_at, updated_at
		FROM offers
		WHERE status = true
		ORDER BY RANDOM()
		LIMIT 25`

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
