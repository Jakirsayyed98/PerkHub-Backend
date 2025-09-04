package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"database/sql"
	"errors"
)

type TicketStore struct {
	db *sql.DB
}

func NewTicketStore(dbs *sql.DB) *TicketStore {
	return &TicketStore{
		db: dbs,
	}
}

func (s *TicketStore) CreateTicket(req *request.CreateTicketRequest, userId string) (string, error) {
	request := model.NewTicket()
	err := request.Bind(userId, req.Subject)
	if err != nil {
		return "", err
	}
	ticketId, err := model.InsertTicket(s.db, request)
	if err != nil {
		return "", err
	}

	ticketMsg := model.NewTicketMessage(ticketId, "user", req.Body)
	err = model.InsertTicketMessage(s.db, ticketMsg)
	if err != nil {
		return "", err
	}
	return ticketId, nil
}

func (s *TicketStore) SendTicketMessage(req *request.SendTicketMsg, userId string) (string, error) {
	ticketMsg := model.NewTicketMessage(req.TicketId, "user", req.Message)
	err := model.InsertTicketMessage(s.db, ticketMsg)
	if err != nil {
		return "", err
	}
	return ticketMsg.TicketID, nil
}

func (s *TicketStore) AdminSentTicketReply(req *request.AdminReplyTicketMsg, userId string) (string, error) {
	ticketMsg := model.NewTicketMessage(req.TicketId, "admin", req.Message)
	err := model.InsertTicketMessage(s.db, ticketMsg)
	if err != nil {
		return "", err
	}
	return ticketMsg.TicketID, nil
}

func (s *TicketStore) GetTicketsByUserId(userId string) ([]*model.Ticket, error) {
	if userId == "" {
		return nil, errors.New("user ID is required")
	}
	return model.GetTicketsByUserId(s.db, userId)
}

func (s *TicketStore) GetTicketMessagesByTicketID(ticketId string) ([]model.TicketMessage, error) {
	if ticketId == "" {
		return nil, errors.New("ticket ID is required")
	}
	return model.GetTicketMessagesByTicketID(s.db, ticketId)
}

func (s *TicketStore) GetTicketsByStatus(ticketStatus string) ([]*model.Ticket, error) {
	if ticketStatus == "" {
		return nil, errors.New("ticket Status is required")
	}

	result, err := model.GetTicketsByStatus(s.db, ticketStatus)
	if err != nil {
		return nil, err
	}
	return result, nil
}
