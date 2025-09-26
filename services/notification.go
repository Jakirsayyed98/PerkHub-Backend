package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2/google"
)

const fcmEndpointTemplate = "https://fcm.googleapis.com/v1/projects/%s/messages:send"

type NotificationService struct {
	projectID          string
	serviceAccountFile string
}

func NewNotificationService(projectID, serviceAccountFile string) *NotificationService {
	return &NotificationService{
		projectID:          projectID,
		serviceAccountFile: serviceAccountFile,
	}
}

type Message struct {
	Message struct {
		Token        string                 `json:"token,omitempty"`
		Topic        string                 `json:"topic,omitempty"`
		Notification map[string]string      `json:"notification,omitempty"`
		Data         map[string]interface{} `json:"data,omitempty"`
	} `json:"message"`
}

// --- helpers ---

func (s *NotificationService) getAccessToken() (string, error) {
	data, err := os.ReadFile(s.serviceAccountFile)
	if err != nil {
		return "", err
	}

	conf, err := google.JWTConfigFromJSON(data, "https://www.googleapis.com/auth/firebase.messaging")
	if err != nil {
		return "", err
	}

	token, err := conf.TokenSource(context.Background()).Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (s *NotificationService) send(msg Message) error {
	accessToken, err := s.getAccessToken()
	if err != nil {
		return err
	}
	fmt.Println("Access Token:", accessToken)

	url := fmt.Sprintf(fcmEndpointTemplate, s.projectID)
	fmt.Println("FCM URL:", url)

	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read full response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("FCM v1 request failed with status %s: %s", resp.Status, string(respBody))
	}

	// Optional: parse response JSON
	var fcmResp map[string]interface{}
	if err := json.Unmarshal(respBody, &fcmResp); err == nil {
		fmt.Println("Parsed FCM Response:", fcmResp)
	}

	return nil
}

// --- Public APIs ---

func (s *NotificationService) SendNotificationToToken(token, title, message string, data map[string]interface{}) error {
	msg := Message{}
	msg.Message.Token = token
	msg.Message.Notification = map[string]string{"title": title, "body": message}
	msg.Message.Data = data
	return s.send(msg)
}

func (s *NotificationService) SendNotificationToTopic(topic, title, message string, data map[string]interface{}) error {
	msg := Message{}
	msg.Message.Topic = topic
	msg.Message.Notification = map[string]string{"title": title, "body": message}
	msg.Message.Data = data
	return s.send(msg)
}

func (s *NotificationService) SendNotificationToAllUsers(title, message string, data map[string]interface{}) error {
	return s.SendNotificationToTopic("all_users", title, message, data)
}
