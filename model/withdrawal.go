package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type WithdrawalRequest struct {
	ID              uuid.UUID  `db:"id" json:"id"`
	UserID          string     `db:"user_id" json:"user_id"`
	RequestedAmt    float64    `db:"requested_amt" json:"requested_amt"`
	ProcessedAmt    *float64   `db:"processed_amt" json:"processed_amt"`
	PaymentMethodID uuid.UUID  `db:"payment_method_id" json:"payment_method_id"`
	Status          string     `db:"status" json:"status"` // 'pending', 'approved', 'rejected'
	Reason          *string    `db:"reason" json:"reason"`
	AdminID         *uuid.UUID `db:"admin_id" json:"admin_id"`
	TxnID           *string    `db:"txn_id" json:"txn_id"`
	TxnTime         *string    `db:"txn_time" json:"txn_time"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at"`
}

func NewWithdrawalRequest() *WithdrawalRequest {
	return &WithdrawalRequest{}
}

func (wr *WithdrawalRequest) Bind(userId, PaymentMethodID string, requestAmount float64) error {
	paymentModeID, err := uuid.Parse(PaymentMethodID)
	if err != nil {
		return err
	}
	wr.UserID = userId
	wr.RequestedAmt = requestAmount
	wr.PaymentMethodID = paymentModeID
	return nil
}

func InsertWithdrawalRequest(db *sql.DB, wr *WithdrawalRequest) (string, error) {
	query := `
        INSERT INTO withdrawal_requests (user_id, requested_amt, payment_method_id)
VALUES ($1, $2, $3)
        RETURNING id;
    `

	var id string
	err := db.QueryRow(
		query,
		wr.UserID,
		wr.RequestedAmt,
		wr.PaymentMethodID,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetWithdrawalByUser(db *sql.DB, userID string) ([]WithdrawalRequest, error) {
	query := `SELECT id,user_id,requested_amt,processed_amt,payment_method_id, status,reason,txn_id,txn_time,created_at FROM withdrawal_requests WHERE user_id = $1 ORDER BY created_at DESC;`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var withdrawalRequest []WithdrawalRequest
	for rows.Next() {
		var wr WithdrawalRequest
		if err := rows.Scan(
			&wr.ID,
			&wr.UserID,
			&wr.RequestedAmt,
			&wr.ProcessedAmt,
			&wr.PaymentMethodID,
			&wr.Status,
			&wr.Reason,
			&wr.TxnID,
			&wr.TxnTime,
			&wr.CreatedAt,
		); err != nil {
			return nil, err
		}
		withdrawalRequest = append(withdrawalRequest, wr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return withdrawalRequest, nil
}

func GetAdminWithdrawalByStatus(db *sql.DB, status string) ([]WithdrawalRequest, error) {
	query := `SELECT id,user_id,requested_amt,processed_amt,payment_method_id, status,reason,txn_id,txn_time,created_at FROM withdrawal_requests WHERE status = $1 ORDER BY created_at DESC;`
	rows, err := db.Query(query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var withdrawalRequest []WithdrawalRequest
	for rows.Next() {
		var wr WithdrawalRequest
		if err := rows.Scan(
			&wr.ID,
			&wr.UserID,
			&wr.RequestedAmt,
			&wr.ProcessedAmt,
			&wr.PaymentMethodID,
			&wr.Status,
			&wr.Reason,
			&wr.TxnID,
			&wr.TxnTime,
			&wr.CreatedAt,
		); err != nil {
			return nil, err
		}
		withdrawalRequest = append(withdrawalRequest, wr)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return withdrawalRequest, nil
}

func ApproveWithdrawal(db *sql.DB, withdrawalID string, processedAmt float64, adminID uuid.UUID, txnID string) error {
	query := `
        UPDATE withdrawal_requests
        SET status = 'approved',
            processed_amt = $2,
            admin_id = $3,
            txn_id = $4,
            txn_time = NOW(),
            updated_at = NOW()
        WHERE id = $1;
    `

	_, err := db.Exec(query, withdrawalID, processedAmt, adminID, txnID)
	if err != nil {
		return fmt.Errorf("failed to approve withdrawal %s: %w", withdrawalID, err)
	}

	return nil
}

func RejectWithdrawal(db *sql.DB, withdrawalID uuid.UUID, reason string, adminID uuid.UUID) error {
	query := `
	  UPDATE withdrawal_requests SET status = 'rejected',
	  reason = $2,
	  admin_id = $3,
	  updated_at = NOW()
	  WHERE id = $1;
    `

	_, err := db.Exec(query, withdrawalID, reason, adminID)
	if err != nil {
		return fmt.Errorf("failed to reject withdrawal %s: %w", withdrawalID, err)
	}

	return nil
}
