package model

import (
	"PerkHub/constants"
	"PerkHub/request"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type MiniAppTransactionData struct {
	Pending    string                `json:"pending"`
	Verified   string                `json:"verified"`
	Rejected   string                `json:"rejected"`
	Withdrawal string                `json:"withdrawal"`
	Data       []MiniAppTransactions `json:"Data"`
}
type MiniAppTransactions struct {
	Id                   string    `json:"id"`
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
	MiniAppID            string    `json:"miniapp_id"`
	CommissionPercentage string    `json:"commission_percentage"`
	TransactionDate      string    `json:"transaction_date"`
	TransactionID        string    `json:"transaction_id"`
	CreatedAt            string    `json:"created_at"`
	UpdatedAt            string    `json:"updated_at"`
	MiniAppName          string    `json:"miniapp_name"`
	Icon                 string    `db:"icon" json:"icon"`
	MiniApp              []MiniApp `json:"miniApp"`
}

func NewMiniAppTransaction() *MiniAppTransactions {
	return &MiniAppTransactions{}
}

func (s *MiniAppTransactions) Bind(req *request.CueLinkCallBackRequest, userCommission, userId, miniappId string) error {

	s.CampaignID = req.CampaignID
	s.Commission = req.Commission
	s.UserCommission = userCommission
	s.UserId = userId
	s.OrderID = req.OrderID
	s.ReferenceID = req.ReferenceID
	s.SaleAmount = req.SaleAmount
	s.Status = req.Status
	s.SubID = req.SubID
	s.SubID1 = req.TransactionID
	s.SubID2 = req.SubID2
	s.MiniAppID = miniappId
	s.CommissionPercentage = req.CommissionPercentage
	s.TransactionDate = req.TransactionDate
	s.TransactionID = req.TransactionID

	return nil
}

func FindMiniAppTransactionByUserID(db *sql.DB, userId string) ([]MiniAppTransactions, error) {
	query := `
		SELECT mt.id,
       mt.campaign_id,
       mt.commission,
       mt.user_commission,
       mt.user_id,
       mt.order_id,
       mt.reference_id,
       mt.sale_amount,
       mt.status,
       mt.subid,
       mt.subid1,
       mt.subid2,
       mt.miniapp_id,
       mt.commission_percentage,
       mt.transaction_date,
       mt.transaction_id,
       mt.created_at,
       mt.updated_at,
       md.name,
	   md.icon
FROM miniapp_transactions mt
LEFT JOIN miniapp_data md
  ON mt.miniapp_id::uuid = md.id
WHERE mt.user_id = $1
ORDER BY mt.created_at DESC`

	// query := "SELECT id,campaign_id, commission, user_commission,user_id, order_id, reference_id, sale_amount, status,
	// subid, subid1, subid2,miniapp_id, commission_percentage, transaction_date, transaction_id, created_at, updated_at
	//  FROM miniapp_transactions WHERE user_id = $1"

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
			&transaction.MiniAppID,
			&transaction.CommissionPercentage,
			&transaction.TransactionDate,
			&transaction.TransactionID,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.MiniAppName,
			&transaction.Icon,
		)
		if err != nil {
			return nil, err
		}
		// miniApp, err := GetMiniAppByID(db, transaction.MiniAppID)

		// if miniApp != nil {
		// 	transaction.MiniApp = append(transaction.MiniApp, *miniApp)
		// }
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

	err := db.QueryRow(query, req.SubID, req.SubID2, req.OrderID).Scan(
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
		&transaction.MiniAppID,
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
	query := "UPDATE miniapp_transactions SET status = $1 , updated_at = $5  WHERE subid = $2 AND subid1 = $3 AND subid2 = $4"
	status := "pending" // default
	if val, ok := constants.StatusMap[strings.ToLower(req.Status)]; ok {
		status = val
	}

	_, err := db.Exec(query, status, req.SubID, req.SubID2, req.OrderID, time.Now())
	if err != nil {
		fmt.Println("Error updating transaction:", err)
		return err
	}
	return nil
}

func InsertMiniAppTransaction(db *sql.DB, req *MiniAppTransactions) error {
	sqlQuery := `INSERT INTO miniapp_transactions ( campaign_id, commission, user_commission, user_id,order_id, reference_id, sale_amount, status, subid, subid1, subid2,miniapp_id, commission_percentage, transaction_date, transaction_id, created_at, updated_at )
	 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`

	status := "pending" // default
	if val, ok := constants.StatusMap[strings.ToLower(req.Status)]; ok {
		status = val
	}

	_, err := db.Exec(sqlQuery, req.CampaignID, req.Commission, req.UserCommission, req.UserId, req.OrderID, req.ReferenceID, req.SaleAmount, status, req.SubID, req.SubID2, req.OrderID, req.MiniAppID, req.CommissionPercentage, req.TransactionDate, req.TransactionID, time.Now(), time.Now())

	return err
}
func GetAllAffiliateTransactions(db *sql.DB, page, limit int, status string) ([]MiniAppTransactions, error) {
	// Calculate offset
	// 	offset := (page - 1) * limit
	// 	// SQL query with LIMIT and OFFSET for pagination
	// 	query := fmt.Sprintf(`
	// 		SELECT id, campaign_id, commission, user_commission, user_id, order_id, reference_id,
	//        sale_amount, status, subid, subid1, subid2, miniapp_id, commission_percentage,
	//        transaction_date, transaction_id, created_at, updated_at
	// FROM miniapp_transactions
	// ORDER BY created_at DESC
	// LIMIT 10 OFFSET 0
	// `)
	if val, ok := constants.StatusMap[strings.ToLower(status)]; ok {
		status = val
	}
	query := `
		SELECT mt.id,
       mt.campaign_id,
       mt.commission,
       mt.user_commission,
       mt.user_id,
       mt.order_id,
       mt.reference_id,
       mt.sale_amount,
       mt.status,
       mt.subid,
       mt.subid1,
       mt.subid2,
       mt.miniapp_id,
       mt.commission_percentage,
       mt.transaction_date,
       mt.transaction_id,
       mt.created_at,
       mt.updated_at,
       md.name
FROM miniapp_transactions mt
LEFT JOIN miniapp_data md
  ON mt.miniapp_id::uuid = md.id
WHERE mt.status = $1
ORDER BY mt.created_at DESC;

	`
	rows, err := db.Query(query, status)
	if err != nil {
		return nil, err
	}
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
			&transaction.MiniAppID,
			&transaction.CommissionPercentage,
			&transaction.TransactionDate,
			&transaction.TransactionID,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.MiniAppName,
		)
		if err != nil {
			return nil, err
		}
		// Fetch miniapp info as before
		// miniApp, err := GetMiniAppByID(db, transaction.MiniAppID)
		// if err != nil {
		// 	return nil, err
		// }
		// if miniApp != nil {
		// 	transaction.MiniApp = append(transaction.MiniApp, *miniApp)
		// }

		transactions = append(transactions, transaction)
	}

	// Check for any errors after row iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
