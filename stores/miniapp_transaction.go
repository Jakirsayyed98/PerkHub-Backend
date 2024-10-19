package stores

import (
	"PerkHub/model"
	"database/sql"
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

	return transaction, nil

}
