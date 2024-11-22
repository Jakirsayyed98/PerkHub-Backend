package model

import (
	"PerkHub/request"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserCashWithdrawal struct {
	Id            uuid.UUID `json:"id"`
	Requested_Amt string    `json:"requested_amt"`
	UserId        string    `json:"user_id"`
	Reason        string    `json:"reason"`
	VPA_ID        string    `json:"VPA_ID"`
	Status        string    `json:"status"`
	TxnId         string    `json:"txn_id"`
	TxnTime       string    `json:"txn_time"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func NewUserCashWithdrawal() *UserCashWithdrawal {
	return &UserCashWithdrawal{}
}

func InserWithdrawalRequest(sql *sql.DB, req request.WithdrawalRequest, userId string) error {
	query := `INSERT INTO user_cash_withdrawal (requested_amt, VPA_ID,user_id, created_at, updated_at) 
				VALUES ($1, $2,$3, NOW(), NOW());`
	_, err := sql.Exec(query, req.RequestedAmt, req.Upi, userId)
	if err != nil {
		return err
	}
	return nil
}

func WithdrawalTxnList(db *sql.DB, userId string) ([]UserCashWithdrawal, error) {
	var reason sql.NullString
	query := "SELECT id,requested_amt,user_id, reason, VPA_ID,status, created_at, updated_at FROM user_cash_withdrawal WHERE user_id = $1"

	rows, err := db.Query(query, userId)
	defer rows.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}

	var transactions []UserCashWithdrawal

	for rows.Next() {
		var transaction UserCashWithdrawal

		err := rows.Scan(
			&transaction.Id,
			&transaction.Requested_Amt,
			&transaction.UserId,
			&reason,
			&transaction.VPA_ID,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if reason.Valid {
			transaction.Reason = reason.String
		} else {
			transaction.Reason = ""
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
