package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"PerkHub/services"
	"database/sql"
)

type NotificationStore struct {
	db                  *sql.DB
	notificationService *services.NotificationService
}

func NewNotificationStore(dbs *sql.DB) *NotificationStore {
	notificationService := services.NewNotificationService()
	return &NotificationStore{
		db:                  dbs,
		notificationService: notificationService,
	}
}

func (s *NotificationStore) CreateNotification(request *request.NotificationRequest) error {
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
		return err
	}
	return nil
}

func (s *NotificationStore) UpdateNotification(request *request.NotificationRequest) error {
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
		return err
	}
	return nil
}

func (s *NotificationStore) GetAdminNotificationList() (interface{}, error) {
	notifications, err := model.AdminGetAllNotification(s.db)
	if err != nil {
		return nil, err
	}
	return notifications, nil
}
func (s *NotificationStore) AdminSendNotification(id string) (interface{}, error) {
	notifications, err := model.GetNotificationByID(s.db, id)
	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
		"click_action": notifications.ClickAction,
		"event":        notifications.EventType,
	}

	s.notificationService.SendNotificationToAllUsers(notifications.Title, notifications.Message, data)
	return notifications, nil
}
