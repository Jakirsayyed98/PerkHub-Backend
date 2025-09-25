package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Notification represents a scheduled or immediate notification
type Notification struct {
	ID          string    `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Message     string    `db:"message" json:"message"`
	EventType   string    `db:"event_type" json:"event_type"`
	StartDate   string    `db:"start_date" json:"start_date,omitempty"`
	EndDate     string    `db:"end_date" json:"end_date,omitempty"`
	Frequency   string    `db:"frequency" json:"frequency,omitempty"`
	ClickAction string    `db:"click_action" json:"click_action,omitempty"`
	Image       string    `db:"image" json:"image,omitempty"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Status      bool      `db:"status" json:"status"`
}

func NewNotification() *Notification {
	return &Notification{}
}

func InsertNotification(db *sql.DB, n *Notification) error {
	err := db.QueryRow(`
		INSERT INTO notification_master (title, message,image, event_type, start_date, end_date, frequency, click_action, created_at, updated_at, status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW(),NOW(),$9)
		RETURNING id, created_at, updated_at
		`, n.Title, n.Message, n.Image, n.EventType, n.StartDate, n.EndDate, n.Frequency, n.ClickAction, n.Status).
		Scan(&n.ID, &n.CreatedAt, &n.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func UpdateNotification(db *sql.DB, update *Notification) error {
	if update.ID == "" {
		return errors.New("missing Notification ID for update")
	}

	clauses := []string{}
	params := []interface{}{}
	i := 1

	// helper to reduce repetition
	addField := func(name string, value interface{}) {
		clauses = append(clauses, fmt.Sprintf("%s=$%d", name, i))
		params = append(params, value)
		i++
	}

	// only update if non-empty/non-nil
	if update.Title != "" {
		addField("title", update.Title)
	}
	if update.Message != "" {
		addField("message", update.Message)
	}

	// only update if non-empty/non-nil
	if update.EventType != "" {
		addField("event_type", update.EventType)
	}
	if update.StartDate != "" {
		addField("start_date", update.StartDate)
	}

	if update.Frequency != "" {
		addField("frequency", update.Frequency)
	}

	// only update if non-empty/non-nil
	if update.ClickAction != "" {
		addField("click_action", update.ClickAction)
	}
	if update.Image != "" {
		addField("image", update.Image)
	}

	addField("status", update.Status)

	// always update timestamp
	addField("updated_at", time.Now())

	if len(clauses) == 0 {
		return errors.New("no valid fields to update")
	}

	query := "UPDATE notification_master SET " +
		strings.Join(clauses, ", ") +
		fmt.Sprintf(" WHERE id=$%d", i)
	params = append(params, update.ID)

	_, err := db.Exec(query, params...)
	return err
}

func AdminGetAllNotification(db *sql.DB) ([]Notification, error) {
	query := "SELECT id,title,message,image,event_type,start_date,end_date,frequency,click_action,status, created_at, updated_at from notification_master"
	result, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var notifications []Notification

	for result.Next() {
		var notification Notification
		err := result.Scan(
			&notification.ID,
			&notification.Title,
			&notification.Message,
			&notification.Image,
			&notification.EventType,
			&notification.StartDate,
			&notification.EndDate,
			&notification.Frequency,
			&notification.ClickAction,
			&notification.Status,
			&notification.CreatedAt,
			&notification.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		notifications = append(notifications, notification)
	}
	if err = result.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func GetNotificationByID(db *sql.DB, id string) (*Notification, error) {
	var notification Notification
	query := "SELECT id,title,message,image,event_type,start_date,end_date,frequency,click_action,status, created_at, updated_at from notification_master WHERE id =$1"
	err := db.QueryRow(query, id).Scan(
		&notification.ID,
		&notification.Title,
		&notification.Message,
		&notification.Image,
		&notification.EventType,
		&notification.StartDate,
		&notification.EndDate,
		&notification.Frequency,
		&notification.ClickAction,
		&notification.Status,
		&notification.CreatedAt,
		&notification.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &notification, nil
}
