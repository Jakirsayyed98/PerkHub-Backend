package stores

import (
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/request"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"
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
	startTime := time.Now()
	paymentMethod := model.NewAddPaymentMethodRequest()
	if err := paymentMethod.Bind(userId, req.Upi, req.BankAccountNumber, req.IFSCCode, req.BankName); err != nil {
		log := logger.LogData{
			Message:   fmt.Sprintf("Error While binding Payment Methods %s ", err.Error()),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	id, err := model.InsertPaymentMethod(s.db, paymentMethod)
	if err != nil {
		log := logger.LogData{
			Message:   fmt.Sprintf("Error While Add Payment Methods %s ", err.Error()),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return id, nil
}

func (s *WithdrawalStore) RequestWithdrawal(req *request.WithdrawalRequest, userId string) (interface{}, error) {
	startTime := time.Now()
	amt, _ := strconv.ParseFloat(req.RequestedAmt, 64)
	if amt < 100 {
		log := logger.LogData{
			Message:   "amount should be greater then 100",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("amount should be greater then 100")
	}

	withdrawal := model.NewWithdrawalRequest()
	if err := withdrawal.Bind(userId, req.PaymentMethod, amt); err != nil {
		return nil, err
	}

	withdrawalCount, err := model.PendingWithdrawalCount(s.db, userId)
	if err != nil {
		log := logger.LogData{
			Message:   fmt.Sprintf("Error While Fetching Pending Withdrawal Count %s ", err.Error()),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	if withdrawalCount >= 1 {
		log := logger.LogData{
			Message:   "You have a pending withdrawal request. Please wait for it to be processed before making a new request.",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("You have a pending withdrawal request. Please wait for it to be processed before making a new request.")
	}

	id, err := model.InsertWithdrawalRequest(s.db, withdrawal)
	if err != nil {
		log := logger.LogData{
			Message:   fmt.Sprintf("Error While Insert Request Payment %s ", err.Error()),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return id, nil
}

func (s *WithdrawalStore) WithdrawalTxnList(userId string) (interface{}, error) {
	startTime := time.Now()
	result, err := model.GetWithdrawalByUser(s.db, userId)
	if err != nil {
		log := logger.LogData{
			Message:   fmt.Sprintf("Error While Get Payment Request List %s ", err.Error()),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return result, nil

}

func (s *WithdrawalStore) AdminWithdrawalTxnList(status string) (interface{}, error) {
	startTime := time.Now()
	result, err := model.GetAdminWithdrawalByStatus(s.db, status)
	if err != nil {
		log := logger.LogData{
			Message:   fmt.Sprintf("Error While Admin Get Payment Requests %s ", err.Error()),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return result, nil

}

func (s *WithdrawalStore) GetPaymentMethodsByUserID(userId string) (interface{}, error) {
	startTime := time.Now()
	if userId == "" {
		log := logger.LogData{
			Message:   "Invalid Data",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("Invalid request")
	}
	result, err := model.GetPaymentMethodsByUserID(s.db, userId)
	if err != nil {
		log := logger.LogData{
			Message:   fmt.Sprintf("Error While Get Payment Methods %s ", err.Error()),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return result, nil
}
