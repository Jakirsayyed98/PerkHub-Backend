package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"database/sql"
	"errors"
	"strconv"
)

type WithdrawalStore struct {
	db *sql.DB
}

func NewWithdrawalStore(dbs *sql.DB) *WithdrawalStore {
	return &WithdrawalStore{
		db: dbs,
	}
}

func (s *WithdrawalStore) RequestWithdrawal(req *request.WithdrawalRequest, userId string) (interface{}, error) {
	amt, _ := strconv.ParseFloat(req.RequestedAmt, 64)
	if amt < 100 {
		return nil, errors.New("amount should be greater then 100")
	}
	if err := model.InserWithdrawalRequest(s.db, *req, userId); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *WithdrawalStore) WithdrawalTxnList(userId string) (interface{}, error) {
	result, err := model.WithdrawalTxnList(s.db, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
