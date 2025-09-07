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

func (s *WithdrawalStore) AddPaymentMethod(req *request.AddPaymentMethodRequest, userId string) (interface{}, error) {
	paymentMethod := model.NewAddPaymentMethodRequest()
	if err := paymentMethod.Bind(userId, req.Upi, req.BankAccountNumber, req.IFSCCode, req.BankName); err != nil {
		return nil, err
	}
	id, err := model.InsertPaymentMethod(s.db, paymentMethod)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *WithdrawalStore) RequestWithdrawal(req *request.WithdrawalRequest, userId string) (interface{}, error) {
	amt, _ := strconv.ParseFloat(req.RequestedAmt, 64)
	if amt < 100 {
		return nil, errors.New("amount should be greater then 100")
	}

	withdrawal := model.NewWithdrawalRequest()
	if err := withdrawal.Bind(userId, req.PaymentMethod, amt); err != nil {
		return nil, err
	}

	id, err := model.InsertWithdrawalRequest(s.db, withdrawal)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *WithdrawalStore) WithdrawalTxnList(userId string) (interface{}, error) {
	result, err := model.GetWithdrawalByUser(s.db, userId)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (s *WithdrawalStore) AdminWithdrawalTxnList(status string) (interface{}, error) {
	result, err := model.GetAdminWithdrawalByStatus(s.db, status)
	if err != nil {
		return nil, err
	}
	return result, nil

}

func (s *WithdrawalStore) GetPaymentMethodsByUserID(userId string) (interface{}, error) {
	if userId == "" {
		return nil, errors.New("Invalid request")
	}
	result, err := model.GetPaymentMethodsByUserID(s.db, userId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
