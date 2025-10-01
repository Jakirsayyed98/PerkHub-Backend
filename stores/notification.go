package stores

import (
	"PerkHub/constants"
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/request"
	"PerkHub/services"
	"database/sql"
	"time"
)

type NotificationStore struct {
	db                  *sql.DB
	notificationService *services.NotificationService
}

func NewNotificationStore(dbs *sql.DB) *NotificationStore {
	notificationService := services.NewNotificationService(
		constants.FirebaseProjectID, constants.FireBaseFilePath,
	)
	return &NotificationStore{
		db:                  dbs,
		notificationService: notificationService,
	}
}

func (s *NotificationStore) CreateNotification(request *request.NotificationRequest) error {
	startTime := time.Now()
	err := model.InsertNotification(s.db, &model.Notification{
		Title:       request.Title,
		Message:     request.Message,
		EventType:   request.EventType,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Frequency:   request.Frequency,
		ClickAction: request.ClickAction,
		Image:       request.Image,
		Status:      request.Status,
	})
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return err
	}
	return nil
}

func (s *NotificationStore) UpdateNotification(request *request.NotificationRequest) error {
	startTime := time.Now()
	err := model.UpdateNotification(s.db, &model.Notification{
		ID:          request.Id,
		Title:       request.Title,
		Message:     request.Message,
		EventType:   request.EventType,
		StartDate:   request.StartDate,
		EndDate:     request.EndDate,
		Frequency:   request.Frequency,
		ClickAction: request.ClickAction,
		Image:       request.Image,
		Status:      request.Status,
	})
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return err
	}
	return nil
}

func (s *NotificationStore) GetAdminNotificationList() (interface{}, error) {
	startTime := time.Now()
	notifications, err := model.AdminGetAllNotification(s.db)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return notifications, nil
}
func (s *NotificationStore) AdminSendNotification(id string) (interface{}, error) {
	startTime := time.Now()
	notifications, err := model.GetNotificationByID(s.db, id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	data := map[string]interface{}{
		"click_action": notifications.ClickAction,
		"event":        notifications.EventType,
		"image":        notifications.Image,
	}

	err = s.notificationService.SendNotificationToAllUsers(notifications.Title, notifications.Message, data)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	err = model.InsertUserNotificationHistory(s.db, &model.UserNotificationHistory{
		Title:       notifications.Title,
		Message:     notifications.Message,
		Image:       notifications.Image,
		ClickAction: notifications.ClickAction,
		UserID:      "",
	})
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

func (s *NotificationStore) GetNotificationById(id string) (interface{}, error) {
	startTime := time.Now()
	notifications, err := model.GetNotificationByID(s.db, id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return notifications, nil
}

func (s *NotificationStore) GetUserNotificationHistory(userId string) (interface{}, error) {
	startTime := time.Now()
	notifications, err := model.GetUserNotificationHistory(s.db, userId)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return notifications, nil
}

func (s *NotificationStore) SendNotificationToUser(userId, title, message, icon, clickAction string) (interface{}, error) {
	startTime := time.Now()
	user, err := model.UserDetailByUserID(s.db, userId)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	data := map[string]interface{}{
		"click_action": clickAction,
		"event":        "shopping",
		"image":        icon,
	}
	err = s.notificationService.SendNotificationToToken(user.FCMToken.String, title, message, data)

	err = model.InsertUserNotificationHistory(s.db, &model.UserNotificationHistory{
		Title:       title,
		Message:     message,
		Image:       icon,
		ClickAction: clickAction,
		UserID:      userId,
	})
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
