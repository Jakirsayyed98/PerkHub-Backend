package stores

import (
	"PerkHub/constants"
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/request"
	"PerkHub/responses"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type AffiliatesStore struct {
	db *sql.DB
}

func NewAffiliatesStore(dbs *sql.DB) *AffiliatesStore {
	return &AffiliatesStore{
		db: dbs,
	}
}

func (s *AffiliatesStore) CueLinkCallBack(req *request.CueLinkCallBackRequest) (interface{}, error) {
	startTime := time.Now()
	UserCommisionPercentage := constants.GET_USER_COMMISION_PERCENTAGE
	if req.Commission == "" {
		req.Commission = "0"
	}
	commission, err := strconv.ParseFloat(req.Commission, 64)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	userCommissionPercentageInt, err := strconv.Atoi(UserCommisionPercentage)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	subIdData, err := model.GetDatabySubId(s.db, req.SubID, req.SubID2)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	usercommision := (commission / 100) * float64(userCommissionPercentageInt)

	request := model.NewMiniAppTransaction()
	if err := request.Bind(req, fmt.Sprintf("%.2f", usercommision), subIdData.UserID, subIdData.StoreID); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	request.CommissionPercentage = UserCommisionPercentage

	miniAppData, err := model.GetMiniAppByID(s.db, subIdData.StoreID)
	if err != nil {
		return nil, err
	}

	_, err = model.FindMiniAppTransactionBySubID(s.db, req)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		if err.Error() != "transaction not found" {
			log := logger.LogData{
				Message:   err.Error(),
				StartTime: startTime,
			}
			logger.LogError(log)
			return nil, err
		}

		err = model.InsertMiniAppTransaction(s.db, request)
		if err != nil {
			log := logger.LogData{
				Message:   err.Error(),
				StartTime: startTime,
			}
			logger.LogError(log)
			return nil, err
		}

		s.CashbackTrackedNotification(subIdData.UserID, miniAppData.Name, fmt.Sprintf("%.2f", usercommision), req.Status, miniAppData.Icon)

		return nil, nil
	}
	err = model.UpdateMiniAppTransaction(s.db, req)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	s.CashbackTrackedNotification(subIdData.UserID, miniAppData.Name, fmt.Sprintf("%.2f", usercommision), req.Status, miniAppData.Icon)
	return nil, nil
}

func (s *AffiliatesStore) CashbackTrackedNotification(userID, storeName, userCommission, status, icon string) error {
	var title, message string

	switch strings.ToLower(status) {
	case "pending":
		title = fmt.Sprintf("💰 Cashback Tracked at %s", storeName)
		message = fmt.Sprintf("🎉 You’ve earned ₹%s from %s! Your cashback has been tracked successfully. 🚀", userCommission, storeName)

	case "validated":
		title = fmt.Sprintf("💰 Cashback Validated at %s", storeName)
		message = fmt.Sprintf("🎉 Good news! Your cashback of ₹%s from %s has been validated. Keep shopping to earn more! 💳", userCommission, storeName)

	case "payable":
		title = fmt.Sprintf("💰 Cashback Verified at %s", storeName)
		message = fmt.Sprintf("✅ Your cashback of ₹%s from %s has been verified and is on its way! 🎊", userCommission, storeName)

	case "rejected":
		title = fmt.Sprintf("❌ Cashback Rejected at %s", storeName)
		message = fmt.Sprintf("⚠️ Unfortunately, your cashback of ₹%s from %s has been rejected. Don’t worry—shop again and grab your rewards! 💳✨", userCommission, storeName)

	default:
		title = fmt.Sprintf("ℹ️ Cashback Update at %s", storeName)
		message = fmt.Sprintf("Your cashback of ₹%s from %s has an update: %s", userCommission, storeName, status)
	}

	// Send notification
	NewNotificationStore(s.db).SendNotificationToUser(
		userID,
		title,
		message,
		icon,
		"order_history",
	)
	return nil
}

func (s *AffiliatesStore) GetCashBackStatusFromAffiliate(req *request.CueLinkCallBackRequest) (*responses.Transaction, error) {

	return nil, nil
}

func (s *AffiliatesStore) CreateAffiliate(req *request.CreateAffiliateRequest) (interface{}, error) {
	startTime := time.Now()
	// Create a new affiliate in the database
	err := model.CreateAffiliate(s.db, req)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) UpdateAffiliate(req *request.CreateAffiliateRequest) (interface{}, error) {
	startTime := time.Now()
	id, err := uuid.Parse(req.Id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, fmt.Errorf("invalid affiliate ID: %w", err)
	}

	// Update the affiliate in the database
	err = model.UpdateAffiliate(s.db, req, id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) DeleteAffiliate(id string) (interface{}, error) {
	startTime := time.Now()
	ids, err := uuid.Parse(id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, fmt.Errorf("invalid affiliate ID: %w", err)
	}
	// Delete the affiliate from the database
	err = model.DeleteAffiliate(s.db, ids)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) UpdateAffiliateFlag(id string, status bool) (interface{}, error) {
	startTime := time.Now()
	ids, err := uuid.Parse(id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, fmt.Errorf("invalid affiliate ID: %w", err)
	}
	// Update the affiliate flag in the database
	err = model.UpdateAffiliateFlag(s.db, ids, status)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return nil, nil
}

func (s *AffiliatesStore) ListAffiliates() (interface{}, error) {
	startTime := time.Now()
	affiliates, err := model.ListAffiliates(s.db)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return affiliates, nil
}

func (s *AffiliatesStore) GetAffiliateByID(id string) (interface{}, error) {
	startTime := time.Now()
	affiliate, err := model.GetAffiliateByID(s.db, id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return affiliate, nil
}
