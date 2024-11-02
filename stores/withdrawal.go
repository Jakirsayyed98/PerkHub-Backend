package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"database/sql"
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
