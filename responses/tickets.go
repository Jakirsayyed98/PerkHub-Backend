package responses

import (
	"PerkHub/model"
	"time"

	"github.com/google/uuid"
)

type TicketResponse struct {
	ID             uuid.UUID        `json:"id"`
	UserID         string           `json:"user_id"`
	Subject        string           `json:"subject"`
	Status         string           `json:"status"`
	CreatedAt      time.Time        `json:"created_at"`
	TicketMessages []*TicketMessage `json:"ticket_messages,omitempty"`
}

type TicketMessage struct {
	ID         uuid.UUID `json:"id"`
	TicketID   string    `json:"ticket_id"`
	AuthorType string    `json:"author_type"` // "user" or "admin"
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created_at"`
}

func NewTicketResponse() *TicketResponse {
	return &TicketResponse{}
}

func (t *TicketResponse) Bind(model *model.Ticket, ticketMessages []model.TicketMessage) error {
	t.ID = uuid.New()
	t.UserID = model.UserID
	t.Subject = model.Subject
	t.Status = model.Status
	t.CreatedAt = model.CreatedAt
	t.TicketMessages = make([]*TicketMessage, len(ticketMessages))
	for i, msg := range ticketMessages {
		t.TicketMessages[i] = &TicketMessage{
			ID:         uuid.New(),
			TicketID:   t.ID.String(),
			AuthorType: msg.AuthorType,
			Body:       msg.Body,
			CreatedAt:  msg.CreatedAt,
		}
	}
	return nil
}
