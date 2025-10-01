package stores

import (
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/request"
	"PerkHub/responses"
	"database/sql"
	"errors"
	"time"
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
	startTime := time.Now()
	request := model.NewTicket()
	err := request.Bind(userId, req.Subject)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return "", err
	}
	ticketId, err := model.InsertTicket(s.db, request)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return "", err
	}

	ticketMsg := model.NewTicketMessage(ticketId, "user", req.Body)
	err = model.InsertTicketMessage(s.db, ticketMsg)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return "", err
	}
	return ticketId, nil
}

func (s *TicketStore) SendTicketMessage(req *request.SendTicketMsg, userId string) (string, error) {
	startTime := time.Now()
	ticketMsg := model.NewTicketMessage(req.TicketId, "user", req.Message)
	err := model.InsertTicketMessage(s.db, ticketMsg)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return "", err
	}
	return ticketMsg.TicketID, nil
}

func (s *TicketStore) AdminSentTicketReply(req *request.AdminReplyTicketMsg, userId string) (string, error) {
	startTime := time.Now()
	ticketMsg := model.NewTicketMessage(req.TicketId, "admin", req.Message)
	err := model.InsertTicketMessage(s.db, ticketMsg)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return "", err
	}
	return ticketMsg.TicketID, nil
}

func (s *TicketStore) ChangeTicketStatus(req *request.ChangeTicketStatusRequest) error {
	startTime := time.Now()
	if req.TicketId == "" {
		log := logger.LogData{
			Message:   "ticket id is required",
			StartTime: startTime,
		}
		logger.LogError(log)
		return errors.New("ticket ID is required")
	}
	if req.Status != "open" && req.Status != "closed" {
		log := logger.LogData{
			Message:   "invalid ticket status",
			StartTime: startTime,
		}
		logger.LogError(log)
		return errors.New("invalid status value")
	}
	return model.UpdateTicketStatus(s.db, req.TicketId, req.Status)
}

func (s *TicketStore) GetTicketsByUserId(userId string) ([]*model.Ticket, error) {
	startTime := time.Now()
	if userId == "" {
		log := logger.LogData{
			Message:   "user id not found",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("user ID is required")
	}
	return model.GetTicketsByUserId(s.db, userId)
}

func (s *TicketStore) GetTicketMessagesByTicketID(ticketId string) (interface{}, error) {
	startTime := time.Now()
	if ticketId == "" {
		log := logger.LogData{
			Message:   "ticket id is required",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("ticket ID is required")
	}

	ticket, err := model.GetTicketByTicketID(s.db, ticketId)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	if ticket == nil {
		log := logger.LogData{
			Message:   "ticket not found",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("ticket not found")
	}

	messages, err := model.GetTicketMessagesByTicketID(s.db, ticketId)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	response := responses.NewTicketResponse()
	err = response.Bind(ticket, messages)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return response, nil
}

func (s *TicketStore) GetTicketsByStatus(ticketStatus string) ([]*model.Ticket, error) {
	startTime := time.Now()
	if ticketStatus == "" {
		log := logger.LogData{
			Message:   "ticket status is required",
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, errors.New("ticket Status is required")
	}

	result, err := model.GetTicketsByStatus(s.db, ticketStatus)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return result, nil
}
