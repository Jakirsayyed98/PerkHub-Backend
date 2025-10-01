package stores

import (
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MiniAppTransactionStore struct {
	db *sql.DB
}

func NewMiniAppTransactionStore(dbs *sql.DB) *MiniAppTransactionStore {
	return &MiniAppTransactionStore{
		db: dbs,
	}
}

func (s *MiniAppTransactionStore) GetMiniAppTransaction(userId string) (interface{}, error) {
	startTime := time.Now()
	transaction, err := model.FindMiniAppTransactionByUserID(s.db, userId)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	pending := 0.0
	verified := 0.0
	rejected := 0.0
	withdrawal := 0.0

	for _, v := range transaction {
		switch v.Status {
		case "pending":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			pending += amt
		case "verified":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			verified += amt
		case "rejected":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			rejected += amt
		case "3":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			withdrawal += amt
		}

	}

	withdrawalList, err := model.GetWithdrawalByUser(s.db, userId)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	for _, v := range withdrawalList {
		amt := v.RequestedAmt
		withdrawal += amt
	}
	return gin.H{
		"pending":    fmt.Sprintf("%.2f", pending),
		"verified":   fmt.Sprintf("%.2f", (verified - withdrawal)),
		"rejected":   fmt.Sprintf("%.2f", rejected),
		"withdrawal": fmt.Sprintf("%.2f", withdrawal),
		"status":     http.StatusOK,
		"message":    "Transaction list",
		"data":       transaction,
	}, nil

}
