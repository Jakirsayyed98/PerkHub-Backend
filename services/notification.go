package services

import (
	"PerkHub/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type NotificationService struct {
	service *settings.HttpService
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		service: settings.NewHttpService("https://fcm.googleapis.com/fcm/send"),
	}
}

const fcmURL = "https://fcm.googleapis.com/fcm/send"

// NotificationPayload is the body we send to FCM
type NotificationPayload struct {
	To              string                 `json:"to,omitempty"`
	RegistrationIDs []string               `json:"registration_ids,omitempty"`
	Notification    map[string]string      `json:"notification"`
	Data            map[string]interface{} `json:"data,omitempty"`
}

// SendNotificationToToken → Single user by token
func (s *NotificationService) SendNotificationToToken(token string, title, message string, data map[string]interface{}) error {
	payload := NotificationPayload{
		To: token,
		Notification: map[string]string{
			"title": title,
			"body":  message,
		},
		Data: data,
	}
	return s.sendFCMRequest(payload)
}

// SendNotificationToTopic → All users subscribed to a topic
func (s *NotificationService) SendNotificationToTopic(topic, title, message string, data map[string]interface{}) error {
	payload := NotificationPayload{
		To: "/topics/" + topic,
		Notification: map[string]string{
			"title": title,
			"body":  message,
		},
		Data: data,
	}
	return s.sendFCMRequest(payload)
}

// SendNotificationToAllUsers → broadcast using "all_users" topic
func (s *NotificationService) SendNotificationToAllUsers(title, message string, data map[string]interface{}) error {
	return s.SendNotificationToTopic("all_users", title, message, data)
}

// internal function to send request
func (s *NotificationService) sendFCMRequest(payload NotificationPayload) error {
	serverKey := os.Getenv("FCM_SERVER_KEY")
	if serverKey == "" {
		return fmt.Errorf("FCM_SERVER_KEY not set")
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fcmURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key="+serverKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("FCM request failed with status: %s", resp.Status)
	}

	return nil
}
