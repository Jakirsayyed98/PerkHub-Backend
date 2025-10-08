package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID              uuid.UUID `json:"id"`
	UserID          string    `json:"user_id"`
	Subject         string    `json:"subject"`
	Status          string    `json:"status"`
	Priority        string    `json:"priority"`
	Category        string    `json:"category"`
	LastMessage     string    `json:"last_message"`
	LastMessageTime time.Time `json:"last_message_time"`
	CreatedAt       time.Time `json:"created_at"`
}

func NewTicket() *Ticket {
	return &Ticket{}
}

func (t *Ticket) Bind(userId, subject, priority, category string) error {
	t.UserID = userId
	t.Subject = subject
	t.Priority = priority
	t.Category = category
	return nil
}

type TicketMessage struct {
	ID         uuid.UUID `json:"id"`
	TicketID   string    `json:"ticket_id"`
	AuthorType string    `json:"author_type"` // "user" or "admin"
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewTicketMessage(ticketID, authorType, body string) *TicketMessage {
	return &TicketMessage{
		TicketID:   ticketID,
		AuthorType: authorType,
		Body:       body,
	}
}

func InsertTicket(db *sql.DB, ticket *Ticket) (string, error) {
	// Implementation for inserting a ticket into the database
	query := "INSERT INTO tickets (user_id, subject, status, priority, category) VALUES ($1, $2, $3, $4, $5) RETURNING id;"
	var newID string
	err := db.QueryRow(query, ticket.UserID, ticket.Subject, "open", ticket.Priority, ticket.Category).Scan(&newID)
	if err != nil {
		return "", err
	}
	return newID, nil
}

func InsertTicketMessage(db *sql.DB, msg *TicketMessage, status string) error {
	query := "INSERT INTO ticket_messages (ticket_id, author_type, body) VALUES ($1, $2, $3);"
	query2 := "UPDATE tickets SET updated_at = NOW(),status=$1 WHERE id = $2;"
	_, err := db.Exec(query, msg.TicketID, msg.AuthorType, msg.Body)
	if err != nil {
		return err
	}
	_, err = db.Exec(query2, status, msg.TicketID)
	return err
}
func GetTicketsByUserId(db *sql.DB, userId string) ([]*Ticket, error) {
	query := `
		SELECT 
			t.id,
			t.user_id,
			t.subject,
			t.status,
			t.priority,
			t.category,
			t.created_at,
			tm.body AS last_message,
			tm.created_at AS last_message_time
		FROM tickets t
		LEFT JOIN LATERAL (
			SELECT 
				body,
				created_at
			FROM ticket_messages
			WHERE ticket_id = t.id
			ORDER BY created_at DESC
			LIMIT 1
		) tm ON true
		WHERE t.user_id = $1
		ORDER BY t.created_at DESC;
	`

	rows, err := db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*Ticket
	for rows.Next() {
		ticket := NewTicket()
		if err := rows.Scan(
			&ticket.ID,
			&ticket.UserID,
			&ticket.Subject,
			&ticket.Status,
			&ticket.Priority,
			&ticket.Category,
			&ticket.CreatedAt,
			&ticket.LastMessage,
			&ticket.LastMessageTime,
		); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func GetTicketByTicketID(db *sql.DB, ticketId string) (*Ticket, error) {
	query := `
			SELECT id, user_id, subject, status, created_at
			FROM tickets
			WHERE id = $1;`

	var ticket Ticket
	err := db.QueryRow(query, ticketId).Scan(&ticket.ID, &ticket.UserID, &ticket.Subject, &ticket.Status, &ticket.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func UpdateTicketStatus(db *sql.DB, ticketId string, status string) error {
	query := `
			UPDATE tickets
			SET status = $1,  updated_at = NOW()
			WHERE id = $2;`
	_, err := db.Exec(query, status, ticketId)
	return err
}

func GetTicketMessagesByTicketID(db *sql.DB, ticketId string) ([]TicketMessage, error) {
	query := `
		SELECT id, ticket_id, author_type, body, created_at
		FROM ticket_messages
		WHERE ticket_id = $1
		ORDER BY created_at ASC;`

	rows, err := db.Query(query, ticketId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []TicketMessage
	for rows.Next() {
		var msg TicketMessage
		if err := rows.Scan(&msg.ID, &msg.TicketID, &msg.AuthorType, &msg.Body, &msg.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	// check for errors after iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func GetTicketsByStatus(db *sql.DB, ticketStatus string) ([]*Ticket, error) {
	query := "SELECT id, user_id, subject, status, created_at FROM tickets WHERE status = $1;"
	rows, err := db.Query(query, ticketStatus)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []*Ticket
	for rows.Next() {
		ticket := NewTicket()
		if err := rows.Scan(&ticket.ID, &ticket.UserID, &ticket.Subject, &ticket.Status, &ticket.CreatedAt); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
