package model

import (
	"PerkHub/request"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MiniAppTransactionData struct {
	Pending    string                `json:"pending"`
	Verified   string                `json:"verified"`
	Rejected   string                `json:"rejected"`
	Withdrawal string                `json:"withdrawal"`
	Data       []MiniAppTransactions `json:"Data"`
}
type MiniAppTransactions struct {
	Id                   uuid.UUID `json:"id"`
	CampaignID           string    `json:"campaign_id"`
	Commission           string    `json:"commission"`
	UserCommission       string    `json:"user_commission"`
	UserId               string    `json:"user_id"`
	OrderID              string    `json:"order_id"`
	ReferenceID          string    `json:"reference_id"`
	SaleAmount           string    `json:"sale_amount"`
	Status               string    `json:"status"`
	SubID                string    `json:"subid"`
	SubID1               string    `json:"subid1"`
	SubID2               string    `json:"subid2"`
	MiniAppId            string    `json:"miniapp_id"`
	CommissionPercentage string    `json:"commission_percentage"`
	TransactionDate      string    `json:"transaction_date"`
	TransactionID        string    `json:"transaction_id"`
	CreatedAt            string    `json:"created_at"`
	UpdatedAt            string    `json:"updated_at"`
}

func NewMiniAppTransaction() *MiniAppTransactions {
	return &MiniAppTransactions{}
}

func (s *MiniAppTransactions) Bind(req *request.CueLinkCallBackRequest, userCommission string) error {

	s.CampaignID = req.CampaignID
	s.Commission = req.Commission
	s.UserCommission = userCommission
	s.UserId = req.SubID2
	s.OrderID = req.ReferenceID
	s.ReferenceID = req.ReferenceID
	s.SaleAmount = req.SaleAmount
	s.Status = req.Status
	s.SubID = req.SubID
	s.SubID1 = req.SubID1
	s.SubID2 = req.SubID2
	s.MiniAppId = req.SubID3
	s.CommissionPercentage = req.CommissionPercentage
	s.TransactionDate = req.TransactionDate
	s.TransactionID = req.TransactionID

	return nil
}

func FindMiniAppTransactionByUserID(db *sql.DB, userId string) ([]MiniAppTransactions, error) {
	query := "SELECT id,campaign_id, commission, user_commission,user_id, order_id, reference_id, sale_amount, status, subid, subid1, subid2,miniapp_id, commission_percentage, transaction_date, transaction_id, created_at, updated_at FROM miniapp_transactions WHERE user_id = $1"

	rows, err := db.Query(query, userId)

	defer rows.Close()

	var transactions []MiniAppTransactions

	for rows.Next() {
		var transaction MiniAppTransactions

		err := rows.Scan(
			&transaction.Id,
			&transaction.CampaignID,
			&transaction.Commission,
			&transaction.UserCommission,
			&transaction.UserId,
			&transaction.OrderID,
			&transaction.ReferenceID,
			&transaction.SaleAmount,
			&transaction.Status,
			&transaction.SubID,
			&transaction.SubID1,
			&transaction.SubID2,
			&transaction.MiniAppId,
			&transaction.CommissionPercentage,
			&transaction.TransactionDate,
			&transaction.TransactionID,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}
	return transactions, nil
}

func FindMiniAppTransactionBySubID(db *sql.DB, req *request.CueLinkCallBackRequest) (*MiniAppTransactions, error) {
	query := "SELECT id,campaign_id, commission, user_commission,user_id, order_id, reference_id, sale_amount, status, subid, subid1, subid2,miniapp_id, commission_percentage, transaction_date, transaction_id, created_at, updated_at FROM miniapp_transactions WHERE subid = $1 AND subid1 = $2 AND subid2 = $3"
	transaction := MiniAppTransactions{}

	err := db.QueryRow(query, req.SubID, req.SubID1, req.SubID2).Scan(
		&transaction.Id,
		&transaction.CampaignID,
		&transaction.Commission,
		&transaction.UserCommission,
		&transaction.UserId,
		&transaction.OrderID,
		&transaction.ReferenceID,
		&transaction.SaleAmount,
		&transaction.Status,
		&transaction.SubID,
		&transaction.SubID1,
		&transaction.SubID2,
		&transaction.MiniAppId,
		&transaction.CommissionPercentage,
		&transaction.TransactionDate,
		&transaction.TransactionID,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, err
	}
	return &transaction, nil
}

func UpdateMiniAppTransaction(db *sql.DB, req *request.CueLinkCallBackRequest) error {
	query := "UPDATE miniapp_transactions SET status = $1 AND updated_at = $5  WHERE subid = $2 AND subid1 = $3 AND subid2 = $4"

	status := "0"
	if req.Status == "Pending" || req.Status == "pending" {
		status = "0"
	} else if req.Status == "Payable" || req.Status == "payable" {
		status = "1"
	} else if req.Status == "Validated" || req.Status == "validated" {
		status = "1"
	} else if req.Status == "Rejected" || req.Status == "rejected" {
		status = "2"
	} else {
		status = "0"
	}

	_, err := db.Exec(query, status, req.SubID, req.SubID1, req.SubID2, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func InsertMiniAppTransaction(db *sql.DB, req *MiniAppTransactions) error {
	sqlQuery := `INSERT INTO miniapp_transactions ( campaign_id, commission, user_commission, user_id,order_id, reference_id, sale_amount, status, subid, subid1, subid2,miniapp_id, commission_percentage, transaction_date, transaction_id, created_at, updated_at )
	 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`

	status := "0"
	if req.Status == "Pending" || req.Status == "pending" {
		status = "0"
	} else if req.Status == "Payable" || req.Status == "payable" {
		status = "1"
	} else if req.Status == "Validated" || req.Status == "validated" {
		status = "1"
	} else if req.Status == "Rejected" || req.Status == "rejected" {
		status = "2"
	} else {
		status = "0"
	}
	_, err := db.Exec(sqlQuery, req.CampaignID, req.Commission, req.UserCommission, req.UserId, req.OrderID, req.ReferenceID, req.SaleAmount, status, req.SubID, req.SubID1, req.SubID2, req.MiniAppId, req.CommissionPercentage, req.TransactionDate, req.TransactionID, time.Now(), time.Now())

	return err
}
