package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PaymentMethod struct {
	ID         uuid.UUID `db:"id"`
	UserID     string    `db:"user_id"`
	Type       string    `db:"type"` // 'upi' or 'bank'
	Identifier string    `db:"identifier"`
	BankName   *string   `db:"bank_name"`
	IFSCCode   *string   `db:"ifsc_code"`
	IsDefault  bool      `db:"is_default"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

func NewAddPaymentMethodRequest() *PaymentMethod {
	return &PaymentMethod{}
}

func (PM *PaymentMethod) Bind(userId, upiId, accountNumber, ifsc, bankName string) error {

	PM.UserID = userId
	if accountNumber != "" && ifsc != "" && bankName != "" {
		PM.Type = "bank"
		PM.Identifier = accountNumber
		PM.BankName = &bankName
		PM.IFSCCode = &ifsc
	} else if upiId != "" {
		PM.Type = "upi"
		PM.Identifier = upiId
	} else {
		return sql.ErrNoRows
	}
	return nil
}

func InsertPaymentMethod(db *sql.DB, pm *PaymentMethod) (string, error) {
	query := `
        INSERT INTO payment_methods (user_id, type, identifier, bank_name, ifsc_code, is_default)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `

	var id string
	err := db.QueryRow(
		query,
		pm.UserID,
		pm.Type,
		pm.Identifier,
		pm.BankName,
		pm.IFSCCode,
		pm.IsDefault,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetPaymentMethodsByUserID(db *sql.DB, userID string) ([]PaymentMethod, error) {
	query := `
	  SELECT id, user_id, type, identifier, bank_name, ifsc_code, is_default
	  FROM payment_methods
	  WHERE user_id = $1;
    `
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []PaymentMethod
	for rows.Next() {
		var pm PaymentMethod
		if err := rows.Scan(
			&pm.ID,
			&pm.UserID,
			&pm.Type,
			&pm.Identifier,
			&pm.BankName,
			&pm.IFSCCode,
			&pm.IsDefault,
		); err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, pm)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return paymentMethods, nil
}

func UpdatePaymentMethodDefault(db sql.DB, userID string, pmID uuid.UUID, isDefault bool) error {
	query := `UPDATE payment_methods SET is_default = $1, updated_at = NOW() WHERE id = $2 AND user_id= $3;`

	_, err := db.Exec(query, isDefault, pmID, userID)
	if err != nil {
		return err
	}
	return nil
}
