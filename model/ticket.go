package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	Subject   string    `json:"subject"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTicket() *Ticket {
	return &Ticket{}
}

func (t *Ticket) Bind(userId, subject string) error {
	t.UserID = userId
	t.Subject = subject
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
	query := "INSERT INTO tickets (user_id, subject, status) VALUES ($1, $2, $3) RETURNING id;"
	var newID string
	err := db.QueryRow(query, ticket.UserID, ticket.Subject, "open").Scan(&newID)
	if err != nil {
		return "", err
	}
	return newID, nil
}

func InsertTicketMessage(db *sql.DB, msg *TicketMessage) error {
	query := "INSERT INTO ticket_messages (ticket_id, author_type, body) VALUES ($1, $2, $3);"
	_, err := db.Exec(query, msg.TicketID, msg.AuthorType, msg.Body)
	return err
}

func GetTicketsByUserId(db *sql.DB, userId string) ([]*Ticket, error) {
	query := "SELECT id, user_id, subject, status, created_at FROM tickets WHERE user_id = $1;"
	rows, err := db.Query(query, userId)
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
