package stores

import (
	"PerkHub/model"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

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

	transaction, err := model.FindMiniAppTransactionByUserID(s.db, userId)
	if err != nil {
		return nil, err
	}
	pending := 0.0
	verified := 0.0
	rejected := 0.0
	withdrawal := 0.0

	for _, v := range transaction {
		switch v.Status {
		case "0":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			pending += amt
		case "1":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			verified += amt
		case "2":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			rejected += amt
		case "3":
			amt, _ := strconv.ParseFloat(v.UserCommission, 64)
			withdrawal += amt
		}

	}

	return gin.H{
		"pending":    fmt.Sprintf("%.2f", pending),
		"verified":   fmt.Sprintf("%.2f", verified),
		"rejected":   fmt.Sprintf("%.2f", rejected),
		"withdrawal": fmt.Sprintf("%.2f", withdrawal),

		"status":  http.StatusOK,
		"message": "Transaction list",
		"data":    transaction,
	}, nil

}
