package model

import (
	"database/sql"
)

type UserNotificationHistory struct {
	Id          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Message     string `db:"message" json:"message"`
	Image       string `db:"image" json:"image"`
	ClickAction string `db:"click_action" json:"click_action"`
	Status      bool   `db:"status" json:"status"`
	CreatedAt   string `db:"created_at" json:"created_at"`
	UpdatedAt   string `db:"updated_at" json:"updated_at"`
	UserID      string `db:"user_id" json:"user_id"`
}

func GetUserNotificationHistory(db *sql.DB, userId string) ([]UserNotificationHistory, error) {
	query := `SELECT id,title, message, image,click_action,status, created_at, updated_at, user_id from user_notification_history WHERE user_id=$1  OR user_id IS NULL OR user_id = '' ORDER BY created_at DESC`
	result, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	var notifications []UserNotificationHistory

	for result.Next() {
		var userNotification UserNotificationHistory
		if err := result.Scan(
			&userNotification.Id,
			&userNotification.Title,
			&userNotification.Message,
			&userNotification.Image,
			&userNotification.ClickAction,
			&userNotification.Status,
			&userNotification.CreatedAt,
			&userNotification.UpdatedAt,
			&userNotification.UserID,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, userNotification)
	}
	if err := result.Err(); err != nil {
		return nil, err
	}
	return notifications, nil
}

func InsertUserNotificationHistory(db *sql.DB, userNotification *UserNotificationHistory) error {
	query := `INSERT INTO user_notification_history (title,message,image,click_action,user_id) VALUES ($1,$2,$3,$4,$5)`
	err := db.QueryRow(query, userNotification.Title, userNotification.Message, userNotification.Image, userNotification.ClickAction, userNotification.UserID)
	if err != nil {
		return err.Err()
	}

	return nil
}
